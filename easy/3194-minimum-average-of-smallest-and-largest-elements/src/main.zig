const std = @import("std");
const root = @import("root.zig");

fn average(nums: []const i32) f32 {
    // @splat(scalar: anytype)
    const n = nums.len;

    var acc: i32 = 0;
    for (nums) |value| acc += value;

    var avg: f32 = @floatFromInt(acc);
    avg = @divExact(avg, 2.0);

    std.debug.print("arr: {any}\n", .{nums}); //> { 1, 2, 3, 4, 5, 6, 7, 8, 9 }
    std.debug.print("len: {any}\n", .{n}); //> 9
    std.debug.print("sum: {any}\n", .{acc}); //> 45
    std.debug.print("avg: {any}\n", .{avg}); //> 2.2e1

    return avg;
}

// [See also](https://zig.guide/standard-library/sorting/)
fn minimum_average(nums: []const i32) !f32 {
    var s: []i32 = try std.heap.page_allocator.alloc(i32, nums.len); //create a mutable copy of input array
    defer std.heap.page_allocator.free(s);
    for (nums, 0..) |value, i| s[i] = value; //copy the elements
    std.mem.sort(i32, s, {}, comptime std.sort.asc(i32)); //sort the array in ascending order

    const mid = @divFloor(nums.len, 2);
    var acc = s[0] + s[s.len - 1];
    for (s[0..mid], 0..) |_, i| {
        acc = @min(acc, s[i] + s[s.len - 1 - i]);
    }
    const res: f32 = @floatFromInt(acc);
    return res * 0.5;
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
        try stdout.print(
            "root.add({0}, {1}) -> '{2}'\n",
            .{ 1, 2, root.add(1, 2) },
        );
    }

    { //3194-minimum-average-of-smallest-and-largest-elements
        const nums = .{ 5, 6, 7, 1, 2, 3, 4, 8, 9 };
        std.debug.assert(3 == root.len_satisfy_fn(i32, &nums, is_even));

        const avg = average(&nums);
        std.debug.print("{any}\n", .{avg});
        const res = try minimum_average(&nums);
        std.debug.print("{any}\n", .{res});
        //Output:
        //
        //  min_avg: 5e0
        //  arr: { 5, 6, 7, 1, 2, 3, 4, 8, 9 }
        //  midpoint: 4, len: 9
        //  l: { 1, 2, 3, 4 }, r: { 9, 8, 7, 6 }
    }

    {
        const nums = .{ 5, 6, 7, 1, 2, 3, 4, 8, 9 };
        std.debug.print("nums: {any}\n", .{nums});
        const is_all_even = root.all(i32, &nums, is_even);
        // const is_all_wrapper: bool = root.wrap_all(root.all(i32, &nums, is_even));
        // std.debug.print("is_all_wrapper: {any}\n", .{is_all_wrapper});
        const is_any_even = (root.any(i32, &nums, is_even) == 1);
        std.debug.print("all: {any}\n", .{is_all_even});
        std.debug.print("any: {any}\n", .{is_any_even});
    }

    try bw.flush(); // don't forget to flush!
}

fn is_even(x: i32) bool {
    return @mod(x, 3) == 0;
}

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
