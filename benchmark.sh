#!/bin/bash
echo "Clean up previous runs"
rm -r logs/*
docker-compose down --rmi=local

COMPOSE_PREFIX="go-rest-barebones_"
APP_CONTAINER_NAME="${COMPOSE_PREFIX}go"
DATABASE_CONTAINER_NAME="${COMPOSE_PREFIX}database"
WEB_CONTAINER_NAME="${COMPOSE_PREFIX}web"

echo "Start Docker Compose Services"
docker-compose up -d

echo "Wait to settle"
sleep 1

echo "Get Container ids"
APP_CONTAINER_ID=$(docker ps -aqf "name=${APP_CONTAINER_NAME}")
DATABASE_CONTAINER_ID=$(docker ps -aqf "name=${DATABASE_CONTAINER_NAME}")
WEB_CONTAINER_NAME=$(docker ps -aqf "name=${WEB_CONTAINER_NAME}")

for i in {1..5}
do
  echo "Start benchmarking run $i/5"
  echo "Start Logging"
  echo "Set up headers in log files"
  echo "CPU%\tMem%\tMemUsage" > logs/app-usage-${i}.log
  echo "CPU%\tMem%\tMemUsage" > logs/db-usage-${i}.log
  echo "CPU%\tMem%\tMemUsage" > logs/web-usage-${i}.log
  # Start App Usage Log
  (timeout 35 docker stats --format '{{.CPUPerc}}\t{{.MemPerc}}\t{{.MemUsage}}' $APP_CONTAINER_ID | sed 's/\x1b\[[0-9;]*[a-zA-Z]//g' ; echo) >> logs/app-usage-${i}.log &
  # Start Database Usage Log
  (timeout 35 docker stats --format '{{.CPUPerc}}\t{{.MemPerc}}\t{{.MemUsage}}' $DATABASE_CONTAINER_ID | sed 's/\x1b\[[0-9;]*[a-zA-Z]//g' ; echo) >> logs/db-usage-${i}.log &
  # Start Web Usage Log
  (timeout 35 docker stats --format '{{.CPUPerc}}\t{{.MemPerc}}\t{{.MemUsage}}' $WEB_CONTAINER_ID | sed 's/\x1b\[[0-9;]*[a-zA-Z]//g' ; echo) >> logs/web-usage-${i}.log &
  # Start benchmark
  echo "Starting benchmark in 5 seconds"
  sleep 5; wrk -d30s -c48 -t8 http://localhost:8888/customer/1 -s ./wrk/report.lua &
  wait
  echo "Benchmark run complete."
done

echo "Completed all benchmarks. See ./logs for results"