<?php

declare(strict_types=1);

namespace App\CsFixer;

use PhpCsFixer\AbstractFixer;
use PhpCsFixer\Tokenizer\FCT;
use PhpCsFixer\Tokenizer\Token;
use PhpCsFixer\Tokenizer\Tokens;
use PhpCsFixer\FixerDefinition\CodeSample;
use PhpCsFixer\FixerDefinition\FixerDefinition;
use PhpCsFixer\Fixer\WhitespacesAwareFixerInterface;
use PhpCsFixer\FixerDefinition\FixerDefinitionInterface;

/**
 * Places one argument per line in method attributes with two or more arguments
 * (e.g. #[Route(path: ..., name: ..., methods: ...)]).
 *
 * Scope is intentionally limited: one attribute with arguments per #[...] group
 * (not grouped #[Foo(...), Bar(...)] attributes). Already multiline groups are
 * left unchanged, making the fixer idempotent.
 */
final class SplitMethodAttributeArgsFixer extends AbstractFixer implements WhitespacesAwareFixerInterface
{
    private const int MIN_ARGS_TO_SPLIT = 2;

    public function getDefinition(): FixerDefinitionInterface
    {
        return new FixerDefinition(
            'Places one argument per line in method attributes with two or more arguments.',
            [
                new CodeSample(
                    <<<'CODE_SAMPLE'
<?php
class SomeController
{
    #[Route(path: '/host/rooms/{id}/edit', name: 'app_host_room_edit', methods: ['GET'])]
    public function edit(int $id): Response
    {
    }
}

CODE_SAMPLE
                ),
            ]
        );
    }

    public function getName(): string
    {
        return '__FIXER_NAMESPACE__/split_method_attribute_args';
    }

    public function isCandidate(Tokens $tokens): bool
    {
        return $tokens->isTokenKindFound(FCT::T_ATTRIBUTE);
    }

    protected function applyFix(\SplFileInfo $file, Tokens $tokens): void
    {
        $groups = [];
        $index = 0;

        while (null !== $index = $tokens->getNextTokenOfKind($index, [[FCT::T_ATTRIBUTE]])) {
            $endIndex = $tokens->findBlockEnd(Tokens::BLOCK_TYPE_ATTRIBUTE, $index);
            $groups[] = [$index, $endIndex];
            $index = $endIndex + 1;
        }

        // Process from right to left: insertions in a group shift only indexes after it,
        // so preceding groups that have not yet been processed remain valid.
        foreach (array_reverse($groups) as [$startIndex, $endIndex]) {
            if (!$this->isMethodAttribute($tokens, $endIndex)) {
                continue;
            }

            $this->fixGroup($tokens, $startIndex, $endIndex);
        }
    }

    private function isMethodAttribute(Tokens $tokens, int $attributeCloseIndex): bool
    {
        $index = $attributeCloseIndex;

        while (true) {
            $index = $tokens->getNextMeaningfulToken($index);

            if (null === $index) {
                return false;
            }

            $token = $tokens[$index];

            if ($token->isGivenKind(\T_FUNCTION)) {
                return true;
            }

            if ($token->isGivenKind(FCT::T_ATTRIBUTE)) {
                $index = $tokens->findBlockEnd(Tokens::BLOCK_TYPE_ATTRIBUTE, $index);
                continue;
            }

            if ($token->isGivenKind([\T_PUBLIC, \T_PROTECTED, \T_PRIVATE, \T_STATIC, \T_FINAL, \T_ABSTRACT, \T_READONLY])) {
                continue;
            }

            return false;
        }
    }

    private function fixGroup(Tokens $tokens, int $startIndex, int $endIndex): void
    {
        $openParenIndex = null;

        for ($i = $startIndex + 1; $i < $endIndex; $i++) {
            $content = $tokens[$i]->getContent();

            if ('(' === $content) {
                $openParenIndex = $i;
                break;
            }

            // A top-level comma before any '(' means multiple bare attributes are grouped
            // (#[Foo, Bar(...)]); this is out of scope, so leave it unchanged.
            if (',' === $content) {
                return;
            }
        }

        if (null === $openParenIndex) {
            return;
        }

        $closeParenIndex = $tokens->findBlockEnd(Tokens::BLOCK_TYPE_PARENTHESIS, $openParenIndex);

        // #[Foo(...), Bar(...)]: multiple attributes in the same group are out of scope.
        $afterClose = $tokens->getNextMeaningfulToken($closeParenIndex);
        if (null !== $afterClose && $afterClose < $endIndex && $tokens[$afterClose]->equals(',')) {
            return;
        }

        $commaIndexes = $this->findTopLevelCommas($tokens, $openParenIndex, $closeParenIndex);

        if (\count($commaIndexes) < self::MIN_ARGS_TO_SPLIT - 1) {
            return;
        }

        if ($this->isAlreadyMultiline($tokens, $openParenIndex, $closeParenIndex)) {
            return;
        }

        $baseIndent = $this->getLineIndent($tokens, $startIndex);
        $argIndent = $baseIndent . $this->whitespacesConfig->getIndent();
        $lineEnding = $this->whitespacesConfig->getLineEnding();

        // Process right to left for the same reason as above.
        $this->setWhitespaceBefore($tokens, $closeParenIndex, $lineEnding . $baseIndent);

        foreach (array_reverse($commaIndexes) as $commaIndex) {
            $this->setWhitespaceAfter($tokens, $commaIndex, $lineEnding . $argIndent);
        }

        $this->setWhitespaceAfter($tokens, $openParenIndex, $lineEnding . $argIndent);
    }

    /**
     * @return int[]
     */
    private function findTopLevelCommas(Tokens $tokens, int $openIndex, int $closeIndex): array
    {
        $commas = [];
        $depth = 0;

        for ($i = $openIndex + 1; $i < $closeIndex; $i++) {
            $content = $tokens[$i]->getContent();

            // Compare raw content rather than Token::equals(): php-cs-fixer assigns a
            // custom kind to literal array '[' and ']' tokens (CT::T_ARRAY_BRACKET_*),
            // so equals('[') does not match them even though they delimit a nesting level.
            if (\in_array($content, ['(', '[', '{'], true)) {
                $depth++;
                continue;
            }

            if (\in_array($content, [')', ']', '}'], true)) {
                $depth--;
                continue;
            }

            if (0 === $depth && ',' === $content) {
                $commas[] = $i;
            }
        }

        return $commas;
    }

    private function isAlreadyMultiline(Tokens $tokens, int $openIndex, int $closeIndex): bool
    {
        for ($i = $openIndex; $i <= $closeIndex; $i++) {
            if (str_contains($tokens[$i]->getContent(), "\n")) {
                return true;
            }
        }

        return false;
    }

    private function getLineIndent(Tokens $tokens, int $attributeStartIndex): string
    {
        $prevIndex = $attributeStartIndex - 1;

        if ($prevIndex < 0 || !$tokens[$prevIndex]->isWhitespace()) {
            return '';
        }

        $content = $tokens[$prevIndex]->getContent();
        $lastNewlinePos = strrpos($content, "\n");

        if (false === $lastNewlinePos) {
            return '';
        }

        return substr($content, $lastNewlinePos + 1);
    }

    private function setWhitespaceAfter(Tokens $tokens, int $index, string $whitespace): void
    {
        $nextIndex = $index + 1;

        if (isset($tokens[$nextIndex]) && $tokens[$nextIndex]->isWhitespace()) {
            $tokens[$nextIndex] = new Token([\T_WHITESPACE, $whitespace]);

            return;
        }

        $tokens->insertAt($nextIndex, new Token([\T_WHITESPACE, $whitespace]));
    }

    private function setWhitespaceBefore(Tokens $tokens, int $index, string $whitespace): void
    {
        $prevIndex = $index - 1;

        if (isset($tokens[$prevIndex]) && $tokens[$prevIndex]->isWhitespace()) {
            $tokens[$prevIndex] = new Token([\T_WHITESPACE, $whitespace]);

            return;
        }

        $tokens->insertAt($index, new Token([\T_WHITESPACE, $whitespace]));
    }
}
