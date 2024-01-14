import { useEffect, useState } from "react";
import { Post, User } from "../types";
import { PostCard } from "./PostCard";

type PostParams = {
    post: Post
}

const DEFAULT_USER: User = {
    avatar: "",
    createdAt: "",
    description: "",
    email: "",
    fullName: "",
    gender: "",
    id: "",
    username: ""
}

export function Post({ post }: PostParams) {
    const [user, setUser] = useState<User>(DEFAULT_USER)
    // TODO: create useUsers hook
    // TODO: create countLikes service
    useEffect(() => {
        const getUser = async () => {
            const res = await fetch("http://localhost:5000/v1/users/devoranico2")
            const data = await res.json()
            setUser(data)
        }

        getUser()
    }, [])
    return (
        <PostCard post={post} user={user} likes={10} />
    )
}