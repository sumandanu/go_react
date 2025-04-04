import { Outlet } from 'react-router-dom';

import styles from './Layout.module.css';
import Navbar from '../Navbar';

const Layout = () => {
  return (
    <div className={styles.container}>
      <Navbar />
      <main>
        <Outlet />
      </main>
      <footer>
        <hr />
        <center>
          Join us for{' '}
          <a href="https://events.ycombinator.com/ai-sus">
            <u>AI Startup School</u>
          </a>{' '}
          this June 16-17 in San Francisco!
        </center>
        <center>
          <span className="yclinks">
            <a href="/newsguidelines">Guidelines</a> | <a href="/newsfaq">FAQ</a> |{' '}
            <a href="lists">Lists</a> | <a href="https://github.com/HackerNews/API">API</a> |{' '}
            <a href="/security">Security</a> |{' '}
            <a href="https://www.ycombinator.com/legal/">Legal</a> |{' '}
            <a href="https://www.ycombinator.com/apply/">Apply to YC</a> |{' '}
            <a href="mailto:hn@ycombinator.com">Contact</a>
          </span>
          <br />
          <br />
          <form method="get" action="//hn.algolia.com/">
            Search:{' '}
            <input
              type="text"
              name="q"
              size={17}
              autoCorrect="off"
              spellCheck="false"
              autoCapitalize="off"
              autoComplete="off"
            />
          </form>
        </center>
      </footer>
    </div>
  );
};

export default Layout;
