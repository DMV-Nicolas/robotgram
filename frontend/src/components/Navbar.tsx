import { Link } from 'react-router-dom'
import { Blank, Create, EmptyHeart, Explore, Github, Home, Search } from './Icons'
import './Navbar.css'

export function Navbar() {
  return (
    <nav className="navbar">
      <h1 className="title">Robotgram</h1>
      <ul>
        <div>
          <li>
            <Link to="/">
              <Home />
              <p>Home</p>
            </Link>
          </li>
          <li>
            <Link to="/search">
              <Search />
              <p>Search</p>
            </Link>
          </li>
          <li>
            <Link to="/explore">
              <Explore />
              <p>Explore</p>
            </Link>
          </li>
          <li>
            <Link to="/notifications">
              <EmptyHeart />
              <p>Notifications</p>
            </Link>
          </li>
          <li>
            <Link to="/create">
              <Create />
              <p>Create</p>
            </Link>
          </li>
          <li>
            <Link to="/profile">
              <img className="avatar" src="https://avatars.githubusercontent.com/u/69326361?v=4" alt="avatar" />
              <p>Profile</p>
            </Link>
          </li>
        </div>
        <div>
          <li>
            <Link to="https://github.com/DMV-Nicolas/robotgram" target="_blank" rel="noreferrer">
              <Github />
              <p>Github</p>
              <span className="navbarLinkBlank">
                <Blank />
              </span>
            </Link>
          </li>
        </div>
      </ul>
    </nav>
  )
}
