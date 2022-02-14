# Lottery
## This is a test assignment that a company my friend works for uses in order to assess the level of Go candidates. Being curious to give it a try, I've decided to write my iwn solution to the task described hereunder:

##Here is the tasks:
###Problem

Imagine a popular free national e-lottery website. In order to take a part in the lottery, users should request a lottery ticket using a button. 

But the number of tickets is limited so only the fastest users who click the button can get it. Each user can get only one ticket.

#The Lottery service

Implement Lottery API service. It should be a simple HTTP server that responds to POST /ticket:
- with HTTP 200 if the request is successful;
- with HTTP 410 if out of tickets;
- with HTTP 403 if users try to issue a ticket again.

**Note:** Provisioning a database is not required for the live technical interview though might be considered in a real environment.

#Requirements:

- users should get the response as fast as possible;
- give out exactly 200 000 tickets;
- each user can only get one ticket;
- means to test the system with ~200k unique users;

**Note:** please, provide additional context and comments regarding the implementation as a separate message or as corresponding documentation along with the home test assignment.
