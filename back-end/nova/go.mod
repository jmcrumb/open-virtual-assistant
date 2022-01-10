module github.com/jmcrumb/nova

go 1.17

replace github.com/jmcrumb/nova/data => ../data

replace github.com/jmcrumb/nova/nlp => ../nlp

require (
	github.com/jmcrumb/nova/data v0.0.0-00010101000000-000000000000
	github.com/jmcrumb/nova/nlp v0.0.0-00010101000000-000000000000
)
