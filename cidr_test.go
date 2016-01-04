package cidr

import (
    "testing"
    "math/rand"
    "net"
)

func TestRange(t *testing.T) {
    r, err := NewRange("127.0.0.0/30")
    if err != nil {
        t.Fatal(err)
    }

    ips := []string{
        "127.0.0.0",
        "127.0.0.1",
        "127.0.0.2",
        "127.0.0.3",
    }

    i := 0
    for {
        if r.String() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
        }

        i++
        if !r.Next() { break; }
    }
}

func TestRangeWith31Prefix(t *testing.T) {
    r, err := NewRangeWithBlockSize("127.0.0.0/30", 31)
    if err != nil {
        t.Fatal(err)
    }

    ips := []string{
        "127.0.0.0",
        "127.0.0.2",
    }

    i := 0
    for i < len(ips)  {
        if r.String() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
        }

        i++
        r.Next()
    }
}

func TestRangeWith25Prefix(t *testing.T) {
    r, err := NewRangeWithBlockSize("127.0.0.0/23", 25)
    if err != nil {
        t.Fatal(err)
    }

    ips := []string{
        "127.0.0.0",
        "127.0.0.128",
        "127.0.1.0",
        "127.0.1.128",
    }

    i := 0
    for i < len(ips) {
        if r.String() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
        }

        i++
        r.Next()
    }
}


func TestRangeWith24Prefix(t *testing.T) {
    r, err := NewRangeWithBlockSize("127.0.0.0/22", 24)
    if err != nil {
        t.Fatal(err)
    }

    ips := []string{
        "127.0.0.0",
        "127.0.1.0",
        "127.0.2.0",
        "127.0.3.0",
    }

    i := 0
    for i < len(ips) {
        if r.String() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
        }

        i++
        r.Next()
    }
}

func TestRangeWith24PrefixShowingPrefix(t *testing.T) {
    r, err := NewRangeWithBlockSize("44.44.0.0/16", 24)
    if err != nil {
        t.Fatal(err)
    }

    ips := []string{
        "44.44.0.0/24",
        "44.44.1.0/24",
        "44.44.2.0/24",
        "44.44.3.0/24",
    }

    i := 0
    for i < len(ips) {
        if r.StringPrefix() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.StringPrefix(), ips[i])
        }

        i++
        r.Next()
    }
}

func TestShouldFailOnBadCidr(t *testing.T) {
    ipranges := []string{
        "127.0.0.1/31",
        "127.0.1.2/30",
        "127.0.2.127/25",
        "127.0.3.129/25",
    }

    for _, iprange := range ipranges {
        _, err := NewRange(iprange)
        if err != nil && err.Error() == "Invalid cidr" {
            // pass
        } else {
            t.Fatal("didn't fail on %s", iprange)
        }
    }
}

func TestShouldFailOnBadBlockStep(t *testing.T) {
    ipranges := []string{
        "127.0.0.0/31",
        "127.0.1.0/30",
        "127.0.2.0/25",
        "127.0.3.0/25",
    }

    for _, iprange := range ipranges {
        _, err := NewRangeWithBlockSize(iprange, 24)
        if err != nil && err.Error() == "Invalid block size" {
            // pass
        } else {
            t.Fatalf("didn't fail on %s", iprange)
        }
    }
}

func TestIP2Long(t *testing.T) {
    testsNum := 32
    for testsNum > 0 {
        ip := net.IPv4(byte(rand.Intn(256)),
            byte(rand.Intn(256)),
            byte(rand.Intn(256)),
            byte(rand.Intn(256)))
        assertIP := Long2IP(IP2Long(ip))
        if !ip.Equal(assertIP) {
            t.Fatalf("didn't work %s != %s", ip.String(), assertIP.String())
            return
        }

        testsNum--
    }
}

func TestLong2IP(t *testing.T) {
    testsNum := 32
    for testsNum > 0 {
        long := uint(rand.Uint32())
        assertLong := IP2Long(Long2IP(long))
        if long != assertLong {
            t.Fatalf("didn't work %u != %u", long, assertLong)
            return
        }

        testsNum--
    }
}
