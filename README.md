# Minimum Railway Platforms

> Striver SDE Sheet 044 · Day 8 Greedy — sort arrivals and departures independently; the sweep's peak occupancy is the platform count. Go, stdlib only, tested.

## Problem

Given every train's arrival and departure time at a station, find the minimum number of platforms so that no train ever waits. A train arriving at the exact minute another departs still needs its own platform (`<=` in the sweep).

## Approach

The pairing of arrival to departure per train is a red herring — only the *count* of trains present matters. Sort arrivals and departures as two independent streams and merge-sweep them chronologically: an arrival increments occupancy, a departure decrements it, and the maximum occupancy ever reached is the answer. This "decouple the events and sweep" move solves a whole family of interval-load problems (meeting rooms II, CPU load, hotel bookings).

**Complexity:** O(n log n) time, O(n) space for the sorted copies.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
