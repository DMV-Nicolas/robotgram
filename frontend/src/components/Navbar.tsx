import { Create, EmptyHeart, Explore, Home, Search } from "./Icons";
import "./Navbar.css"

type NavbarLiItemParams = {
    icon: JSX.Element
    text: string
}

function NavbarLiItem({ icon }: NavbarLiItemParams) {

}

export function Navbar() {
    return (
        <nav className="navbar">
            <h1 className="title">Robotgram</h1>
            <ul>
                <li>
                    <a href="/"><Home /> <p>Home</p></a>
                </li>
                <li>
                    <a href="/search"><Search /> <p>Search</p></a>
                </li>
                <li>
                    <a href="/explore"><Explore /> <p>Explore</p></a>
                </li>
                <li>
                    <a href="/notifications"><EmptyHeart /> <p>Notifications</p></a>
                </li>
                <li>
                    <a href="/create"><Create /> <p>Create</p></a>
                </li>
                <li>
                    <a href="/profile"><img className="avatar" src="https://avatars.githubusercontent.com/u/69326361?v=4" alt="avatar" /> Profile</a>
                </li>
            </ul>
        </nav>
    )
}