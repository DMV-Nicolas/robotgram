export type PostType = {
    id: string
    userID: string
    images: string[]
    description: string
    createdAt: string
}

export type UserType = {
    id: string
    username: string
    fullName: string
    email: string
    avatar: string
    description: string
    gender: string
    createdAt: string
}

export type PostResponse = {
    id: string
    user_id: string
    images: string[]
    description: string
    created_at: string
}

export type UserResponse = {
    id: string
    username: string
    full_name: string
    email: string
    avatar: string
    description: string
    gender: string
    created_at: string
}