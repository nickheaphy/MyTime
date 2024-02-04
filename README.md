# MyTime

For many years I have noted down what I do (work related) on a calendar that I print out every week. This was really for my own informaiton, so I could actually remember what I did - it was especially useful at review time to be able to say what I do.

A few years ago work wanted to see where our time was spent and had us enter our time into ServiceDesk, then this changed to entering time in cases in Salesforce. Neither of these systems were great at being able to quickly translate my written calendar into logs. I did write some tools to make this easier (XXXXXXX) but it still didn't really work the way I wanted to work.

Late last year it was decided that we didn't need to track our time using any of the electronic systems. Problem was, I quite liked the reporting that could be pulled from having all my time logs in a database.

I looked at some existing tools, but nothing appealed. The all looked like logging time would become a chore.

So I decided to hack together something basic, that worked the same way as my paper system, that was quick to translate the paper entries to digital at the end of the week.

## What I want to capture

I didn't want to have to spend a lot of time entering data - but I also wanted to be able to see where my time was being spent. I needed a couple of classifications, some notes, who the work was for, just basic things. I ended up with this list:

- Time/Duration
- Primary Log Type (ie Presales, Project, Postsales, Leave etc)
- Secondary Log Type (ie Presales-Solution Design, Presales-CustomerMeeting etc)
- Text Description
- Customer

## Usability

It had to be simple - I wanted to avoid manual data entry as much as possible (a previous system I used forced all time to be entered in seconds - who thought that was a good idea?)

As I end up working on the same thing over multiple days, I wanted to be able to add new entries but reuse previously entered data so I needed to be able to pull historical data from the database when adding new time logs.

I wanted some sort of colour coding (which is still a bit of a work in progress) so I could see different work types.

I needed some way to generate some reports, with selectable time frames.

I wanted the log classifications to be constrained, and ideally with only a few different options (this goal may not have been met...)

## Installation (Mac)

The tool is written in Go and presents as a web page.

Install Go from https://go.dev/ 

Uses the the CGO free version of sqlite `go get modernc.org/sqlite@latest`

Need SQLite3 to be installed `brew install sqlite3`

## First run

Edit the `dbSQLcmd.go` file to set the categories and colours for each category.

Run using `go run .`

## Notes

This is really quite an ugly hack. It only does very minimal data validation, it has no concept of different users, there is not a GUI for setting things like the event classifications. Really minimal viable product stuff - but it is just for personal use, so didn't want to invest a huge amount of time (though as I am not a great Go or Javascript programmer, it did take way longer that I planned to even get to this stage)

## Todo

- Update the default colour scheme

https://earthly.dev/blog/golang-sqlite/
https://www.w3schools.com/xml/ajax_xmlhttprequest_send.asp
https://forum.golangbridge.org/t/using-golang-to-make-a-string-and-uses-ajax/22652/2
https://medium.com/@edwardpie/processing-form-request-data-in-golang-2dff4c2441be
