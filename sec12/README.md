# Flyweight Design Pattern
A space optimization technique that let us use less memory by storing extenmally the data
associated with similar objects.

## Motivation
### Avoid dedundancy when storing data
### Gaming Apps
* Storing users with identical fist/last names
* Store a list of names and references to them (indices, pointers, etc.)
* No sense in storing same list of first/last name over and over again
### Text Formating
* Don't want each character to have a formattin character 
* Operate on range (e.g., line number, start/end positions)
