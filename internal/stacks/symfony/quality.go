package symfony

import _ "embed"

//go:embed quality/SplitMethodAttributeArgsFixer.php
var splitMethodAttributeArgsFixer string

//go:embed quality/BlankLineAfterControlStructureFixer.php
var blankLineAfterControlStructureFixer string

const qualityPHPStan = `includes:
    - vendor/phpstan/phpstan-doctrine/extension.neon
parameters:
    fileExtensions:
        - php
    level: 6
    paths:
        - bin/
        - config/
        - public/
        - src/
    excludePaths:
        - vendor
`

const qualityRector = `<?php

declare(strict_types=1);

use Rector\Config\RectorConfig;

return RectorConfig::configure()
    ->withPaths([
        __DIR__ . '/bin',
        __DIR__ . '/config',
        __DIR__ . '/public',
        __DIR__ . '/src',
        __DIR__ . '/tests',
    ])
    ->withSkip([
        __DIR__ . '/vendor',
        __DIR__ . '/var',
    ])
    ->withPhpVersion(__PHP_VERSION_ID__)
    ->withPhpSets(php__PHP_MINOR__: true)
    ->withPHPStanConfigs([__DIR__ . '/phpstan.neon'])
    ->withPreparedSets(deadCode: true);
`

const qualityPHPcsFixer = `<?php

declare(strict_types=1);

use App\CsFixer\BlankLineAfterControlStructureFixer;
use App\CsFixer\SplitMethodAttributeArgsFixer;
use PhpCsFixer\Config;
use PhpCsFixer\Finder;

$finder = Finder::create()
    ->in(__DIR__)
    ->exclude([
        '.docker',
        'node_modules',
        'public/build',
        'public/uploads',
        'templates',
        'translations',
        'var',
        'vendor',
    ])
    ->notPath('config/reference.php')
    ->ignoreDotFiles(true)
    ->ignoreVCS(true);

return (new Config())
    ->registerCustomFixers([
        new SplitMethodAttributeArgsFixer(),
        new BlankLineAfterControlStructureFixer(),
    ])
    ->setRules([
        '@PSR12' => true,
        'no_unused_imports' => true,
        'declare_strict_types' => true,
        'array_syntax' => ['syntax' => 'short'],
        'ordered_imports' => ['sort_algorithm' => 'length'],
        'trailing_comma_in_multiline' => true,
        'binary_operator_spaces' => ['default' => 'single_space'],
        'no_extra_blank_lines' => true,
        'no_whitespace_in_blank_line' => true,
        'single_space_around_construct' => true,
        'types_spaces' => ['space' => 'single'],
        'concat_space' => ['spacing' => 'one'],
        'align_multiline_comment' => false,
        'single_quote' => true,
        'class_attributes_separation' => ['elements' => ['method' => 'one']],
        'combine_consecutive_unsets' => true,
        'combine_consecutive_issets' => true,
        'type_declaration_spaces' => true,
        'statement_indentation' => [
            'stick_comment_to_next_continuous_control_statement' => true,
        ],
        '__FIXER_NAMESPACE__/split_method_attribute_args' => true,
        '__FIXER_NAMESPACE__/blank_line_after_control_structure' => true,
    ])
    ->setRiskyAllowed(true)
    ->setUsingCache(true)
    ->setFinder($finder);
`
