

struct Req {
    //
    1: required map<UserKey, i32> users_score
}

struct Resp {
    //
}

struct UserKey {
    1: required i32 id
    2: optional string name
    3: required bool gender
}
