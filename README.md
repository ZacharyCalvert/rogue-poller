# rogue-poller
Rogue Fitness website product availability poller.  For the fitness products you want...but so do 20,000 other people.

# Why?

Why not?  In all seriousness, after buying thousands of dollars of equipment from Rogue Fitness during COVID-19, I identified problems with page caching resulting in missed opportunities and false positives for product availability.  

This project serves to:
- Be my personal pet project for golang
- Build a CLI to notify me when products are available that I would like to purchase

# Short Term

Short print out of what products ar available as a result of running the app, basically crawl the siite.

# Long Term
COVID-19 becomes COVID-22?  Maybe this is integrated into twitter and pages me via email.  

# Running 

Biggest pain is reverse engineering their product pages to find the page and the correct product identity.  

```
go get github.com/ZacharyCalvert/rogue-poller
```
