import { Blank, Create, EmptyHeart, Explore, Github, Home, Search } from './Icons'
import './Navbar.css'

export function Navbar() {
  return (
    <nav className="navbar">
      <h1 className="title">Robotgram</h1>
      <ul>
        <div>
          <li>
            <a href="/">
              <Home />
              <p>Home</p>
            </a>
          </li>
          <li>
            <a href="/search">
              <Search />
              <p>Search</p>
            </a>
          </li>
          <li>
            <a href="/explore">
              <Explore />
              <p>Explore</p>
            </a>
          </li>
          <li>
            <a href="/notifications">
              <EmptyHeart />
              <p>Notifications</p>
            </a>
          </li>
          <li>
            <a href="/create">
              <Create />
              <p>Create</p>
            </a>
          </li>
          <li>
            <a href="/profile">
              <img className="avatar" src="https://avatars.githubusercontent.com/u/69326361?v=4" alt="avatar" />
              <p>Profile</p>
            </a>
          </li>
        </div>
        <div>
          <li>
            <a href="https://github.com/DMV-Nicolas/robotgram" target="_blank" rel="noreferrer">
              <Github />
              <p>Github</p>
              <span className="navbarLinkBlank">
                <Blank />
              </span>
            </a>
          </li>
        </div>
      </ul>
    </nav>
  )
}
