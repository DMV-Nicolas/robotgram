import { Blank, Create, EmptyHeart, Explore, Github, Home, Search } from './Icons'
import './Navbar.css'

export const Navbar = (): JSX.Element => {
  return (
    <nav className="navbar">
      <h1 className="title">Robotgram</h1>
      <ul>
        <div>
          <li>
            <a className="navbarLink" href="/">
              <Home />
              <p>Home</p>
            </a>
          </li>
          <li>
            <a className="navbarLink" href="/search">
              <Search />
              <p>Search</p>
            </a>
          </li>
          <li>
            <a className="navbarLink" href="/explore">
              <Explore />
              <p>Explore</p>
            </a>
          </li>
          <li>
            <a className="navbarLink" href="/notifications">
              <EmptyHeart />
              <p>Notifications</p>
            </a>
          </li>
          <li>
            <a className="navbarLink" href="/create">
              <Create />
              <p>Create</p>
            </a>
          </li>
          <li>
            <a className="navbarLink" href="/profile">
              <img className="avatar" src="https://avatars.githubusercontent.com/u/69326361?v=4" alt="avatar" />
              <p>Profile</p>
            </a>
          </li>
        </div>
        <div>
          <li>
            <a className="navbarLink" href="https://github.com/DMV-Nicolas/robotgram" target="_blank" rel="noreferrer">
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
