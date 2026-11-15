# Pinguin

A unix `ping` analogue built in Go just for fun, no AI, no slop, just vibes & learning.

## TODO notes

- [ ] **Bring direct IP packet handling to project**

  A most of project is implemented, but I faced a problem on getting the `TTL` of an IP packet, so, next obvious step is just to build `ip` packaging to reach at the point, when I can both set & read the `TTL` header from the packet

- [ ] **Packet filtering**

  Kernel by default sends a copy of a received ICMP packet to each process, that waits an ICMP packet, so program can handle "noise" by an accident, need to filter each packet using `identifier`, `seqNum`, `type` & `sourceAddress`.

- [ ] **CLI**

  Current CLI is too primitive, I need an exact copy of `ping` from macOS (Unix?)
