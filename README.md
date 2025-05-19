# go-moon
go implementation of project moonramp 

An offramp to the moon!

Status: Early Concept

## Services
go-moon runs two services:

1. Web Server: Http web server with HTML templates, Tailwind CSS and HTMX. 
2. DB Updater: Makes calls to external API and updates DB on set time interval. DB called from Web Server for up to date information. 

## Introduction
Moonramp allows users to build portfolios of their crypto assets and set cost basis for each asset. Once a portfolio is created, the user can set incremental sell targets based on % increase or target prices. The intent is to allow users to create and manage strategies to Dollar Cost Average (DCA) sell their assets. The Database holds the price vallue for each asset and is updated by the DB Updater service. 