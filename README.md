# GitReporter

this small program in still work-in-progress. The idea is to generate some JSON file that summarize the statistics of a Git repository.

The data are extracted using git command line instructions, than the output is parsed and divided into different hash tables. These are then used to generate one or multiple JSON files, that can be used by external Javascript Plotting library.
