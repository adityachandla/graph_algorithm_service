# This remains constant
runs=20
location="local"
sf=10

algo="bfs"
parallelism=1
go run main.go \
    --parallelism ${parallelism}\
    --repetitions ${runs} \
    --algorithm ${algo} \
    --nodeMap nodeMap${sf}.csv \
    2>> ${location}_${algo}_${parallelism}_${sf}.txt
echo "Finished ${location}_${algo}_${parallelism}_${sf}.txt"

algo="dfs"
parallelism=1
go run main.go \
    --parallelism ${parallelism}\
    --repetitions ${runs} \
    --algorithm ${algo} \
    --nodeMap nodeMap${sf}.csv \
    2>> ${location}_${algo}_${parallelism}_${sf}.txt
echo "Finished ${location}_${algo}_${parallelism}_${sf}.txt"

algo="bfs"
parallelism=2
go run main.go \
    --parallelism ${parallelism}\
    --repetitions ${runs} \
    --algorithm ${algo} \
    --nodeMap nodeMap${sf}.csv \
    2>> ${location}_${algo}_${parallelism}_${sf}.txt
echo "Finished ${location}_${algo}_${parallelism}_${sf}.txt"

algo="dfs"
parallelism=2
go run main.go \
    --parallelism ${parallelism}\
    --repetitions ${runs} \
    --algorithm ${algo} \
    --nodeMap nodeMap${sf}.csv \
    2>> ${location}_${algo}_${parallelism}_${sf}.txt
echo "Finished ${location}_${algo}_${parallelism}_${sf}.txt"
