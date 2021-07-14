import blspy
from chia.types.blockchain_format.sized_bytes import bytes32
from chia.util.ints import uint8, uint64, uint128
from chia.consensus.pos_quality import _expected_plot_size

def std_hash(b) -> bytes32:
    """
    The standard hash used in many places.
    """
    return bytes32(blspy.Util.hash256(bytes(b)))

# 2021-07-14T14:47:40.137 harvester chia.harvester.harvester: INFO     lmd: 2400
# 2021-07-14T14:47:40.138 harvester chia.harvester.harvester: INFO     lmd: 147573952589676412928
# 2021-07-14T14:47:40.138 harvester chia.harvester.harvester: INFO     lmd: 6bc7e96778d56f99640d7d606682543a1449fe84c53d9b3bd764decacc29a10c
# 2021-07-14T14:47:40.139 harvester chia.harvester.harvester: INFO     lmd: 32
# 2021-07-14T14:47:40.139 harvester chia.harvester.harvester: INFO     lmd: 396fef5662016a16c8849b58bcbd6362368792f637b2a7a2abd91db2f35b9a80
# 2021-07-14T14:47:40.140 harvester chia.harvester.harvester: INFO     lmd: Iterations, Required: 1110877509098 SP: 1884160

def main():
    quality_string: bytes32 = bytearray.fromhex("6bc7e96778d56f99640d7d606682543a1449fe84c53d9b3bd764decacc29a10c")
    cc_sp_output_hash: bytes32 = bytearray.fromhex("396fef5662016a16c8849b58bcbd6362368792f637b2a7a2abd91db2f35b9a80")

    sp_quality_string: bytes32 = std_hash(quality_string + cc_sp_output_hash)

    # print(sp_quality_string)

    spa = int.from_bytes(sp_quality_string, "big", signed=False)
    spb = (int(pow(2, 256)) * int(_expected_plot_size(32)))
    sqs =  spa // spb

    print(spa, spb, sqs)

    iters = uint64(
        int(2400)
        * int(147573952589676412928)
        * int.from_bytes(sp_quality_string, "big", signed=False)
        // (int(pow(2, 256)) * int(_expected_plot_size(32)))
    )

    print(iters)
    if (iters == 1110877509098):
        print("MATCH")

if __name__ == "__main__":
    main()
