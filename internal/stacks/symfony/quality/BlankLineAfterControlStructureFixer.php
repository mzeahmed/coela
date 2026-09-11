<?php

declare(strict_types=1);

namespace App\CsFixer;

use PhpCsFixer\AbstractFixer;
use PhpCsFixer\Tokenizer\Token;
use PhpCsFixer\Tokenizer\Tokens;
use PhpCsFixer\FixerDefinition\CodeSample;
use PhpCsFixer\FixerDefinition\FixerDefinition;
use PhpCsFixer\Fixer\WhitespacesAwareFixerInterface;
use PhpCsFixer\FixerDefinition\FixerDefinitionInterface;

/** Adds a blank line after an if, for, or foreach block when followed by another statement. */
final class BlankLineAfterControlStructureFixer extends AbstractFixer implements WhitespacesAwareFixerInterface
{
    private const array CONTROL_TOKENS = [\T_IF, \T_FOR, \T_FOREACH];

    public function getDefinition(): FixerDefinitionInterface
    {
        return new FixerDefinition(
            'Adds a blank line after an if, for, or foreach block when followed by another statement.',
            [
                new CodeSample(
                    <<<'CODE_SAMPLE'
                        <?php
                        if ($condition) {
                            doSomething();
                        }
                        continueExecution();
                        CODE_SAMPLE
                ),
            ],
        );
    }

    public function getName(): string
    {
        return '__FIXER_NAMESPACE__/blank_line_after_control_structure';
    }

    public function isCandidate(Tokens $tokens): bool
    {
        foreach (self::CONTROL_TOKENS as $kind) {
            if ($tokens->isTokenKindFound($kind)) {
                return true;
            }
        }

        return false;
    }

    protected function applyFix(\SplFileInfo $file, Tokens $tokens): void
    {
        $closingBraces = [];

        foreach (self::CONTROL_TOKENS as $kind) {
            $index = 0;
            while (null !== $index = $tokens->getNextTokenOfKind($index, [[$kind]])) {
                $closingBrace = $this->findClosingBrace($tokens, $index);
                if (null !== $closingBrace) {
                    $closingBraces[] = $closingBrace;
                }

                ++$index;
            }
        }

        rsort($closingBraces);
        foreach (array_unique($closingBraces) as $closingBrace) {
            $this->ensureBlankLineAfter($tokens, $closingBrace);
        }
    }

    private function findClosingBrace(Tokens $tokens, int $controlIndex): ?int
    {
        $openParenthesis = $tokens->getNextMeaningfulToken($controlIndex);
        if (null === $openParenthesis || '(' !== $tokens[$openParenthesis]->getContent()) {
            return null;
        }

        $closeParenthesis = $tokens->findBlockEnd(Tokens::BLOCK_TYPE_PARENTHESIS, $openParenthesis);
        $openBrace = $tokens->getNextMeaningfulToken($closeParenthesis);
        if (null === $openBrace || '{' !== $tokens[$openBrace]->getContent()) {
            return null;
        }

        $closingBrace = $tokens->findBlockEnd(Tokens::BLOCK_TYPE_CURLY_BRACE, $openBrace);

        return \T_IF === $tokens[$controlIndex]->getId()
            ? $this->findConditionalChainEnd($tokens, $closingBrace)
            : $closingBrace;
    }

    private function findConditionalChainEnd(Tokens $tokens, int $closingBrace): int
    {
        while (null !== $continuation = $tokens->getNextMeaningfulToken($closingBrace)) {
            if ($tokens[$continuation]->isGivenKind(\T_ELSEIF)) {
                $openParenthesis = $tokens->getNextMeaningfulToken($continuation);
                if (null === $openParenthesis || '(' !== $tokens[$openParenthesis]->getContent()) {
                    return $closingBrace;
                }

                $closeParenthesis = $tokens->findBlockEnd(Tokens::BLOCK_TYPE_PARENTHESIS, $openParenthesis);
                $openBrace = $tokens->getNextMeaningfulToken($closeParenthesis);
            } elseif ($tokens[$continuation]->isGivenKind(\T_ELSE)) {
                $openBrace = $tokens->getNextMeaningfulToken($continuation);
            } else {
                return $closingBrace;
            }

            if (null === $openBrace || '{' !== $tokens[$openBrace]->getContent()) {
                return $closingBrace;
            }

            $closingBrace = $tokens->findBlockEnd(Tokens::BLOCK_TYPE_CURLY_BRACE, $openBrace);
        }

        return $closingBrace;
    }

    private function ensureBlankLineAfter(Tokens $tokens, int $closingBrace): void
    {
        $nextMeaningful = $tokens->getNextMeaningfulToken($closingBrace);
        if (null === $nextMeaningful || $tokens[$nextMeaningful]->equals('}')) {
            return;
        }

        if ($tokens[$nextMeaningful]->isGivenKind([\T_ELSE, \T_ELSEIF])) {
            return;
        }

        $whitespaceIndex = $closingBrace + 1;
        $lineEnding = $this->whitespacesConfig->getLineEnding();
        if (isset($tokens[$whitespaceIndex]) && $tokens[$whitespaceIndex]->isWhitespace()) {
            $whitespace = $tokens[$whitespaceIndex]->getContent();
            if ($this->containsBlankLine($whitespace)) {
                return;
            }

            $tokens[$whitespaceIndex] = new Token([
                \T_WHITESPACE,
                $lineEnding . $lineEnding . $this->indentAfterLastLineEnding($whitespace),
            ]);

            return;
        }

        $tokens->insertAt($whitespaceIndex, new Token([
            \T_WHITESPACE,
            $lineEnding . $lineEnding . $this->lineIndent($tokens, $closingBrace),
        ]));
    }

    private function containsBlankLine(string $whitespace): bool
    {
        return 1 === preg_match('/\R[^\S\r\n]*\R/', $whitespace);
    }

    private function indentAfterLastLineEnding(string $whitespace): string
    {
        $parts = preg_split('/\R/', $whitespace);

        return false === $parts ? '' : (string) end($parts);
    }

    private function lineIndent(Tokens $tokens, int $index): string
    {
        for ($i = $index - 1; $i >= 0; --$i) {
            $content = $tokens[$i]->getContent();
            if (preg_match('/\R([^\r\n]*)$/', $content, $matches)) {
                return $matches[1];
            }
        }

        return '';
    }
}
