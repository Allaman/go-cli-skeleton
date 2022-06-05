all

# Extend line length, since each sentence should be on a separate line.
rule 'MD013', :line_length => 99999
# https://github.com/updownpress/markdown-lint/blob/master/rules/024-no-duplicate-header.md
rule 'MD024', :siblings_only => true
