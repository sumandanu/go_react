import { Link } from 'react-router-dom';
import styles from './Navbar.module.css';
import { useLocation } from 'react-router-dom';

interface IMenu {
  label: string;
  url: string;
}

function Navbar() {
  const location = useLocation();
  const { pathname } = location;

  const menus: IMenu[] = [
    { label: 'new', url: '/newest' },
    { label: 'past', url: '/front' },
    { label: 'comments', url: '/newcomments' },
    { label: 'ask', url: '/ask' },
    { label: 'show', url: '/show' },
    { label: 'jobs', url: '/jobs' },
    { label: 'submit', url: '/submit' }
  ];

  return (
    <nav>
      <div className={styles.logo}>
        <a href="https://news.ycombinator.com">
          <img src="/assets/images/y18.svg" alt="" />
        </a>
      </div>

      <ul>
        <li>
          <Link to="/news" className={styles.main}>
            Hacker News
          </Link>
        </li>
        {menus.map((menu) => (
          <li key={menu.url}>
            <Link to={menu.url} className={`${pathname === menu.url && styles.active}`}>
              {menu.label}
            </Link>
          </li>
        ))}
      </ul>
      <button onClick={() => console.log('')}>login</button>
    </nav>
  );
}

export default Navbar;
