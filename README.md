# DF Exporter

using df to export to prometheus metrics

```bash
~ df
Filesystem     512-blocks      Used Available Capacity iused      ifree %iused  Mounted on
/dev/disk2s7s1  478724992  17729152  59787368    23%  356091  298936840    0%   /
devfs                 420       420         0   100%     728          0  100%   /dev
/dev/disk2s6    478724992   8389128  59787368    13%       4  298936840    0%   /System/Volumes/VM
/dev/disk2s2    478724992  10051936  59787368    15%    1249  298936840    0%   /System/Volumes/Preboot
/dev/disk2s4    478724992    105272  59787368     1%     124  298936840    0%   /System/Volumes/Update
/dev/disk3s2      1024000     12328    987944     2%       4    4939720    0%   /System/Volumes/xarts
/dev/disk3s1      1024000     12704    987944     2%      36    4939720    0%   /System/Volumes/iSCPreboot
```

```
# HELP df_available_space_kilobytes Available space in kilobytes.
# TYPE df_available_space_kilobytes gauge
df_available_space_kilobytes{filesystem="/dev/disk2s1"} 2.9901056e+07
df_available_space_kilobytes{filesystem="/dev/disk2s2"} 2.9901056e+07
df_available_space_kilobytes{filesystem="/dev/disk2s4"} 2.9901056e+07
df_available_space_kilobytes{filesystem="/dev/disk2s5"} 2.9901056e+07
df_available_space_kilobytes{filesystem="/dev/disk2s6"} 2.9901056e+07
df_available_space_kilobytes{filesystem="/dev/disk2s7s1"} 2.9901056e+07
df_available_space_kilobytes{filesystem="/dev/disk3s1"} 493972
df_available_space_kilobytes{filesystem="/dev/disk3s2"} 493972
df_available_space_kilobytes{filesystem="/dev/disk3s3"} 493972
df_available_space_kilobytes{filesystem="/dev/disk4s1"} 146484
df_available_space_kilobytes{filesystem="/dev/disk5s1"} 129480
df_available_space_kilobytes{filesystem="devfs"} 0
df_available_space_kilobytes{filesystem="map"} 0
# HELP df_used_space_kilobytes Used space in kilobytes.
# TYPE df_used_space_kilobytes gauge
df_used_space_kilobytes{filesystem="/dev/disk2s1"} 1.89695864e+08
df_used_space_kilobytes{filesystem="/dev/disk2s2"} 5.025968e+06
df_used_space_kilobytes{filesystem="/dev/disk2s4"} 52636
df_used_space_kilobytes{filesystem="/dev/disk2s5"} 860
df_used_space_kilobytes{filesystem="/dev/disk2s6"} 4.194564e+06
df_used_space_kilobytes{filesystem="/dev/disk2s7s1"} 8.864576e+06
df_used_space_kilobytes{filesystem="/dev/disk3s1"} 6352
df_used_space_kilobytes{filesystem="/dev/disk3s2"} 6164
df_used_space_kilobytes{filesystem="/dev/disk3s3"} 652
df_used_space_kilobytes{filesystem="/dev/disk4s1"} 362684
df_used_space_kilobytes{filesystem="/dev/disk5s1"} 318500
df_used_space_kilobytes{filesystem="devfs"} 210
df_used_space_kilobytes{filesystem="map"} 0
```