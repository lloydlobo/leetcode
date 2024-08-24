const std = @import("std");
const testing = std.testing;

pub fn len_satisfy_fn(comptime T: type, comptime slice: []const T, comptime predicate: fn (T) bool) comptime_int {
    comptime {
        var count: usize = 0;
        for (slice) |value| {
            if (predicate(value))
                count += 1;
        }
        return count;
    }
}

///Returns 1 if all items in `slice` satisfy `predicate`, otherwise 0.
fn _all(comptime T: type, comptime slice: []const T, comptime predicate: fn (T) bool) comptime_int {
    comptime {
        var res: comptime_int = 1; //either 0 or 1
        for (slice) |value| {
            if (res == 0)
                break;
            if (!predicate(value))
                res = 0; //false
        }
        return res;
    }
}

///Returns `true` if all items in `slice` satisfy `predicate`, otherwise `false`.
pub fn all(comptime T: type, comptime slice: []const T, comptime predicate: fn (T) bool) bool {
    return _all(T, slice, predicate) == 1;
}

pub fn wrap_all(comptime maybe0or1: comptime_int) bool {
    return maybe0or1 == 1;
}

///Returns 1 if any item in `slice` satisfy `predicate`, otherwise 0.
pub fn any(comptime T: type, comptime slice: []const T, comptime predicate: fn (T) bool) comptime_int {
    comptime {
        var res = 0; //either 0 or 1
        for (slice) |value| {
            if (res == 1) break;
            if (predicate(value)) res = 1; //true
        }
        return res;
    }
}

///Caller owns and must free the returned slice with all.free(result)
///
///[View in browser](https://www.reddit.com/r/Zig/comments/18z21qs/comment/kghk09r/)
fn map(comptime T: type, comptime S: type, alloc: std.mem.Allocator, xs: []const T, comptime func: fn (T) S) ![]S {
    const res = try alloc.alloc(S, xs.len);
    errdefer alloc.free(res);
    for (xs, res) |in, *out|
        out.* = func(in);
    return res;
}

///Caller owns and must free the returned slice with all.free(result)
///
///[View in browser](https://www.reddit.com/r/Zig/comments/18z21qs/comment/kghk09r/)
fn filter(comptime T: type, alloc: std.mem.Allocator, xs: []const T, comptime prop: fn (T) bool) ![]T {
    var res = try alloc.alloc(T, xs.len);
    errdefer alloc.free(res);
    var i: usize = 0;
    for (xs) |in| {
        if (prop(in)) {
            res[i] = in;
            i += 1;
        }
    }
    const res_resized = try alloc.realloc(res, i);
    return res_resized;
}

// pub fn filter(comptime T: type, comptime slice: []const T, comptime predicate: fn (T) bool) []T {
//
//     //See also:
//     //  - [filter](https://www.reddit.com/r/Zig/comments/18z21qs/comment/kgf6e5k/)
//     comptime {
//         var res: [len_satisfy_fn(T, slice, predicate)]T = undefined;
//         std.debug.print("res initialize: {any}\n", .{res});
//         var i: usize = 0;
//         for (slice) |value| {
//             if (predicate(value)) {
//                 res[i] = value;
//                 i += 1;
//             }
//         }
//         return &res;
//     }
// }

//`export` for C and `pub` for Zig.
pub export fn add(a: i32, b: i32) i32 {
    return a + b;
}

test "basic add functionality" {
    try testing.expect(add(3, 7) == 10);
}
