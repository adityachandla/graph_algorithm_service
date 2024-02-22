# This remains constant
runs=20
location="s3"
sf=10

algo="bfs"
parallelism=1
./algo \
    --parallelism ${parallelism}\
    --repetitions ${runs} \
    --algorithm ${algo} \
    --nodeMap nodeMap${sf}.csv \
    >> ${location}_${algo}_${parallelism}_${sf}.txt 2>&1
echo "Finished ${location}_${algo}_${parallelism}_${sf}.txt"

sleep 10 # These ensure that prefetcher has enough time to catch up.

algo="dfs"
parallelism=1
./algo \
    --parallelism ${parallelism}\
    --repetitions ${runs} \
    --algorithm ${algo} \
    --nodeMap nodeMap${sf}.csv \
    >> ${location}_${algo}_${parallelism}_${sf}.txt 2>&1
echo "Finished ${location}_${algo}_${parallelism}_${sf}.txt"

sleep 10

algo="bfs"
parallelism=2
./algo \
    --parallelism ${parallelism}\
    --repetitions ${runs} \
    --algorithm ${algo} \
    --nodeMap nodeMap${sf}.csv \
    >> ${location}_${algo}_${parallelism}_${sf}.txt 2>&1
echo "Finished ${location}_${algo}_${parallelism}_${sf}.txt"

sleep 10

algo="dfs"
parallelism=2
./algo \
    --parallelism ${parallelism}\
    --repetitions ${runs} \
    --algorithm ${algo} \
    --nodeMap nodeMap${sf}.csv \
    >> ${location}_${algo}_${parallelism}_${sf}.txt 2>&1
echo "Finished ${location}_${algo}_${parallelism}_${sf}.txt"
