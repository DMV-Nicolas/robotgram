export type Post = {
    id: string
    userID: string
    images: string[]
    description: string
    createdAt: string
}

export type User = {
    id: string
    username: string
    fullName: string
    email: string
    avatar: string
    description: string
    gender: string
    createdAt: string
}