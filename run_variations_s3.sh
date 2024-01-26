# This remains constant
runs=100
location="s3"
sf=10

algo="bfs"
parallelism=1
./algo \
    --parallelism ${parallelism}\
    --repetitions ${runs} \
    --algorithm ${algo} \
    --nodeMap nodeMap${sf}.csv \
    2>> ${location}_${algo}_${parallelism}_${sf}.txt
echo "Finished ${location}_${algo}_${parallelism}_${sf}.txt"

algo="dfs"
parallelism=1
./algo \
    --parallelism ${parallelism}\
    --repetitions ${runs} \
    --algorithm ${algo} \
    --nodeMap nodeMap${sf}.csv \
    2>> ${location}_${algo}_${parallelism}_${sf}.txt
echo "Finished ${location}_${algo}_${parallelism}_${sf}.txt"

algo="bfs"
parallelism=2
./algo \
    --parallelism ${parallelism}\
    --repetitions ${runs} \
    --algorithm ${algo} \
    --nodeMap nodeMap${sf}.csv \
    2>> ${location}_${algo}_${parallelism}_${sf}.txt
echo "Finished ${location}_${algo}_${parallelism}_${sf}.txt"

algo="dfs"
parallelism=2
./algo \
    --parallelism ${parallelism}\
    --repetitions ${runs} \
    --algorithm ${algo} \
    --nodeMap nodeMap${sf}.csv \
    2>> ${location}_${algo}_${parallelism}_${sf}.txt
echo "Finished ${location}_${algo}_${parallelism}_${sf}.txt"
