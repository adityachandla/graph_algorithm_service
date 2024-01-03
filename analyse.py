from collections import defaultdict
import re
import sys

def main():
    if len(sys.argv) < 2:
        print("Enter the path to directory")
        exit(0)
    base = sys.argv[1]
    if not base.endswith('/'):
        base += "/"
    server_file = base + "/server.log"
    client_file = base + "/client.log"
    analyze_client(client_file)
    analyze_server(server_file)

def analyze_server(filename):
    print("=========Server Analysis==========")
    with open(filename, "r") as f:
        lines = f.read().strip().split("\n")
    fetching_lines = [l for l in lines if "Fetching" in l and "csr" in l]
    print(f"Total files fetched: {len(fetching_lines)}")
    pat = re.compile(r".* Fetching (.*?).csr")
    dd = defaultdict(int)
    for l in fetching_lines:
        filename = pat.match(l)[1]
        dd[filename] += 1
    print(f"Num distinct files {len(dd)}")
    print("==================================")



def analyze_client(filename):
    print("=========Client Analysis==========")
    with open(filename, "r") as f:
        lines = f.read().strip().split("\n")
    one_hop = [l for l in lines if "OneHop" in l]
    pat = re.compile(r".*? Found (\d+) locations in (\d+)")
    one_hop_times = []
    for l in one_hop:
        res = pat.match(l)
        one_hop_times.append(int(res[2]))
    print(f"One hop average time {sum(one_hop_times)/len(one_hop_times)}ms")
    print(f"One hop cache hits:{cache_hits(one_hop_times)}")
    two_hop = [l for l in lines if "TwoHop" in l]
    pat = re.compile(r".*? Found (\d+) friends of friends in (\d+)")
    two_hop_times = []
    for l in two_hop:
        res = pat.match(l)
        two_hop_times.append(int(res[2]))
    print(f"Two hop average time {sum(two_hop_times)/len(two_hop_times)}ms")
    print(f"Two hop cache hits:{cache_hits(two_hop_times)}")
    pat = re.compile(r".*? Found (\d+) places in (\d+)")
    three_hop = [l for l in lines if "ThreeHop" in l]
    three_hop_times = []
    for l in three_hop:
        res = pat.match(l)
        three_hop_times.append(int(res[2]))
    print(f"Three hop average time {sum(three_hop_times)/len(three_hop_times)}ms")
    print(f"Three hop cache hits:{cache_hits(three_hop_times)}")
    print("==================================")
    print()

def cache_hits(l: list[int]) -> int:
    c = 0
    for v in l:
        if v < 10:
            c += 1
    return c

if __name__ == "__main__":
    main()
