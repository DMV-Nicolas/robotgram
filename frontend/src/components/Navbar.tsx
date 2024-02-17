import { Link } from 'react-router-dom'
import { Blank, Create, EmptyHeart, Explore, Github, Home, Search } from './Icons'
import './Navbar.css'

interface NavbarProps {
  userID: string
}

export function Navbar({ userID }: NavbarProps) {
  return (
    <nav className="navbar">
      <h1 className="navbar__title">Robotgram</h1>
      <ul className='navbar__ul'>
        <div className='navbar__ulGroup'>
          <li className='navbar__item'>
            <Link to="/">
              <Home />
              <p>Home</p>
            </Link>
          </li>
          <li className='navbar__item'>
            <Link to="/search">
              <Search />
              <p>Search</p>
            </Link>
          </li>
          <li className='navbar__item'>
            <Link to="/explore">
              <Explore />
              <p>Explore</p>
            </Link>
          </li>
          <li className='navbar__item'>
            <Link to="/notifications">
              <EmptyHeart />
              <p>Notifications</p>
            </Link>
          </li>
          <li className='navbar__item'>
            <Link to="/create">
              <Create />
              <p>Create</p>
            </Link>
          </li>
          <li className='navbar__item'>
            <Link to={`/user/${userID}`}>
              <img className="avatar" src="https://avatars.githubusercontent.com/u/69326361?v=4" alt="avatar" />
              <p>Profile</p>
            </Link>
          </li>
        </div>
        <div className='navbar__ulGroup'>
          <li className='navbar__item'>
            <Link to="https://github.com/DMV-Nicolas/robotgram" target="_blank" rel="noreferrer">
              <Github />
              <p>Github</p>
              <span className="navbar__blank">
                <Blank />
              </span>
            </Link>
          </li>
        </div>
      </ul>
    </nav>
  )
}
