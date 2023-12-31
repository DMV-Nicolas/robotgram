import { useEffect, useRef, useState } from "react"
import "./PostCard.css"

export type user = {
    _id: string
    username: string
    avatar: string
    gender: "male" | "female"
}

export type post = {
    _id: string
    user_id: string
    images: string[]
    description: string
    created_at: string
}

type sliderParams = {
    postID: string
    images: string[]
}

function Slider({ images, postID }: sliderParams) {
    const listRef = useRef<HTMLUListElement>(null)
    const [currentIndex, setCurrentIndex] = useState(0)


    useEffect(() => {
        const listNode = listRef.current
        let imgNode: Element
        if (listNode != null) {
            imgNode = listNode.querySelectorAll("li > img")[currentIndex]
            imgNode.scrollIntoView({
                behavior: "smooth"
            })
        }
    }, [currentIndex])


    const scrollToImage = (direction: "prev" | "next") => {
        if (direction === 'prev') {
            setCurrentIndex(curr => {
                const isFirstSlide = currentIndex === 0
                return isFirstSlide ? 0 : curr - 1
            })
        } else {
            const isLastSlide = currentIndex === images.length - 1
            if (!isLastSlide) {
                setCurrentIndex(curr => curr + 1)
            }
        }
    }

    const goToSlide = (slideIndex: number) => {
        setCurrentIndex(slideIndex)
    }

    return (
        <div className="rg-postCard-slider">
            <div className="leftArrow" onClick={() => scrollToImage("prev")}></div>
            <div className="rightArrow" onClick={() => scrollToImage("next")}></div>
            <div className="rg-postCard-images">
                <ul ref={listRef}>
                    {images.map((url, i) => <li key={postID + i}>
                        <img className="rg-postCard-image" src={url} width={468} height={585} />
                    </li>)}
                </ul>
            </div>
            <div className="dots-container">
                {images.map((_, idx) => (<div key={idx}
                    className={`dot-container-item ${idx === currentIndex ? "active" : ""}`}
                    onClick={() => goToSlide(idx)}>
                    &#9865;
                </div>))}
            </div>
        </div>
    )

}

type postCardParams = {
    user: user
    post: post
    likes: number
}

export function PostCard({ post, user, likes }: postCardParams) {
    let genderEmoji: string
    if (user.gender == "male") {
        genderEmoji = "♂️"
    } else {
        genderEmoji = "♀️"
    }

    return (
        <article className="rg-postCard">
            <header className="rg-postCard-header">
                <img className="rg-postCard-avatar" src={user.avatar} alt="avatar" />
                <strong className="rg-postCard-username">{user.username} {genderEmoji}</strong>
            </header>
            <div className="rg-postCard-body">
                <Slider images={post.images} postID={post._id} />
            </div>
            <footer className="rg-postCard-footer">
                <span></span>
            </footer>
        </article>
    )
}