# Drone scheduler

[![Build Status](https://drone.notprod.homeoffice.gov.uk/api/badges/UKHomeOffice/dronescheduler/status.svg)](https://drone.notprod.homeoffice.gov.uk/UKHomeOffice/dronescheduler) [![Docker Repository on Quay](https://quay.io/repository/ukhomeofficedigital/dronescheduler/status "Docker Repository on Quay")](https://quay.io/repository/ukhomeofficedigital/dronescheduler)

Currently drone doesn't support scheduled builds. Since we need to build some projects on a regular base, we have created a small container that will run a cron job and will trigger drone builds.

