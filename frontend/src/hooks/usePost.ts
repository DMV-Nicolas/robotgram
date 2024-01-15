import { useState, useEffect } from "react"
import { PostType, UserType, UserResponse } from "../types"

const DEFAULT_USER: UserType = {
    avatar: "",
    createdAt: "",
    description: "",
    email: "",
    fullName: "",
    gender: "",
    id: "",
    username: ""
}

export function usePost({ post }: { post: PostType }) {
    const [user, setUser] = useState(DEFAULT_USER)
    const [likes, setLikes] = useState(0)

    useEffect(() => {
        const getUser = async () => {
            const res = await fetch(`http://localhost:5000/v1/users/${post.userID}`)
            const data: UserResponse = await res.json()
            const user: UserType = {
                id: data.id,
                username: data.username,
                fullName: data.full_name,
                email: data.email,
                avatar: data.avatar,
                description: data.description,
                gender: data.gender,
                createdAt: data.created_at,
            }
            setUser(user)
        }

        getUser()
    }, [post.userID])

    useEffect(() => {
        const getCountOfLikes = async () => {
            const res = await fetch(`http://localhost:5000/v1/likes/${post.id}/count`)
            const data = await res.json()
            setLikes(data)
        }

        getCountOfLikes()
    }, [post.id])

    return { user, likes }
}