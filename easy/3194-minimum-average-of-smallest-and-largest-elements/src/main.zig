const std = @import("std");
const print = std.debug.print;

fn average(arr: []const i32) f32 {
    const n = arr.len;

    var acc: i32 = 0;
    for (arr) |value| acc += value;

    var avg: f32 = @floatFromInt(acc);
    avg = @divExact(avg, 2.0);

    print("arr: {any}\n", .{arr}); //> { 1, 2, 3, 4, 5, 6, 7, 8, 9 }
    print("len: {any}\n", .{n}); //> 9
    print("sum: {any}\n", .{acc}); //> 45
    print("avg: {any}\n", .{avg}); //> 2.2e1

    return avg;
}

fn minimum_average(nums: []const i32) !f32 {
    var s: []i32 = try std.heap.page_allocator.alloc(i32, nums.len); //create a mutable copy of input array
    defer std.heap.page_allocator.free(s);
    for (nums, 0..) |value, i| s[i] = value; //copy the elements
    std.mem.sort(i32, s, {}, comptime std.sort.asc(i32)); //sort the array in ascending order

    const mid = @divFloor(nums.len, 2);

    if (comptime false) {
        const l = s[0..mid]; //create slices for the left and right parts
        var r = try std.heap.page_allocator.alloc(i32, mid);
        defer std.heap.page_allocator.free(r);
        for (l, 0..) |_, i| r[i] = s[s.len - 1 - i]; //reverse the right slice
        var min_sum: i32 = l[0] + r[0]; //initial calculate the minimum sum of pairs from l and r
        for (l, 0..) |_, i| min_sum = @min(min_sum, l[i] + r[i]);
        const res: f32 = @floatFromInt(min_sum);
        return res * 0.5;
    } else {
        var acc = s[0] + s[s.len - 1];
        for (s[0..mid], 0..) |_, i| {
            acc = @min(acc, s[i] + s[s.len - 1 - i]);
        }
        const res: f32 = @floatFromInt(acc);
        return res * 0.5;
    }
}

pub fn main() !void {
    // Prints to stderr (it's a shortcut based on `std.io.getStdErr()`)
    std.debug.print("All your {s} are belong to us.\n", .{"codebase"});

    // stdout is for the actual output of your application, for example if you
    // are implementing gzip, then only the compressed bytes should be sent to
    // stdout, not any debugging messages.
    const stdout_file = std.io.getStdOut().writer();
    var bw = std.io.bufferedWriter(stdout_file);

    const stdout = bw.writer();

    // const allocator = std.heap.page_allocator;
    // var alloc = allocator.alloc(i32, 1024);

    {
        const root = @import("root.zig");
        try stdout.print(
            "root.add({0}, {1}) -> '{2}'\n",
            .{ 1, 2, root.add(1, 2) },
        );
    }

    {
        const nums = .{ 5, 6, 7, 1, 2, 3, 4, 8, 9 };
        const res = try minimum_average(&nums);
        std.debug.print("{any}\n", .{res});
        //Output:
        //
        //  min_avg: 5e0
        //  arr: { 5, 6, 7, 1, 2, 3, 4, 8, 9 }
        //  midpoint: 4, len: 9
        //  l: { 1, 2, 3, 4 }, r: { 9, 8, 7, 6 }
    }

    try bw.flush(); // don't forget to flush!
}

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
