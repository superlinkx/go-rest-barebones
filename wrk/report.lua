done = function(summary, latency, requests)
  path = "./logs/benchmark_results.csv"
  
  -- check if file exists, if not, we want to write a header so set exists to false
  exists = true
  f = io.open(path, "r")
  if f == nil then
    exists = false
  else
    f:close()
  end

  -- open output file
  f = io.open("./logs/benchmark_results.csv", "a+")
  if not exists then
    f:write("min lat, max lat, mean lat, std dev lat, 50 lat, 90 lat, 99 lat, 99.999 lat, dur, tot req, req/s, recv bytes\n")
  end
  
  -- write below results to file
  --   minimum latency
  --   max latency
  --   mean of latency
  --   standard deviation of latency
  --   50percentile latency
  --   90percentile latency
  --   99percentile latency
  --   99.999percentile latency
  --   duration of the benchmark
  --   total requests during the benchmark
  --   requests per second during the benchmark
  --   total received bytes during the benchmark

  f:write(string.format("%f,%f,%f,%f,%f,%f,%f,%f,%d,%d,%f,%d\n",
  latency.min, latency.max, latency.mean, latency.stdev, latency:percentile(50),
  latency:percentile(90), latency:percentile(99), latency:percentile(99.999),
  summary["duration"], summary["requests"], summary["requests"]/summary["duration"]*1000000,
  summary["bytes"]))
  
  f:close()
end
