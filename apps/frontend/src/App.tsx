// import React from 'react';
// import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <>
      <section className="container">
        <ul>
          <li>
            <a href="default.asp">Hacker News</a>
          </li>
          <li>
            <a href="news.asp">new</a>
          </li>
          <li>
            <a href="contact.asp">past</a>
          </li>
          <li>
            <a href="about.asp">comments</a>
          </li>
          <li>
            <a href="about.asp">ask</a>
          </li>
          <li>
            <a href="about.asp">show</a>
          </li>
          <li>
            <a href="about.asp">jobs</a>
          </li>
          <li>
            <a href="about.asp">submit</a>
          </li>
        </ul>
        <ol style={{ display: 'flex', flexDirection: 'column' }}>
          <li>
            <div style={{ display: 'flex', flexDirection: 'column' }}>
              <div style={{ display: 'inline-flex', gap: 10, alignItems: 'end' }}>
                <span>I maintain a 17 year old ThinkPad</span>
                <span style={{ fontSize: 12 }}>( pilledtexts.com )</span>
              </div>
              <div
                style={{
                  display: 'inline-flex',
                  gap: 4,
                  color: 'GrayText',
                  fontSize: 12,
                  alignItems: 'center'
                }}>
                <span>113 points</span>
                <span>by</span>
                <span>Fred34</span>
                <span>3 hours ago</span>
                <span>|</span>
                <span>hide</span>
                <span>|</span>
                <span>79 comments</span>
              </div>
            </div>
          </li>
          <li>test</li>
        </ol>
      </section>

      {/* <ol className="rounded-list">
        <li>
          <a href="">List item</a>
        </li>
        <li>
          <a href="">List item</a>
        </li>
        <li>
          <a href="">List item</a>
          <ol>
            <li>
              <a href="">List sub item</a>
            </li>
            <li>
              <a href="">List sub item</a>
            </li>
            <li>
              <a href="">List sub item</a>
            </li>
          </ol>
        </li>
        <li>
          <a href="">List item</a>
        </li>
        <li>
          <a href="">List item</a>
        </li>
      </ol> */}
    </>
  );
}

export default App;

//  {/* // <div className="App">
//   //   <header className="App-header">
//   //     <img src={logo} className="App-logo" alt="logo" />
//   //     <p>
//   //       Edit <code>src/App.tsx</code> and save to reload.
//   //     </p>
//   //     <a
//   //       className="App-link"
//   //       href="https://reactjs.org"
//   //       target="_blank"
//   //       rel="noopener noreferrer"
//   //     >
//   //       Learn React
//   //     </a>
//   //   </header>
//   // </div> */}
