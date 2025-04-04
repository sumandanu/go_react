import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { FC, ReactNode } from 'react';
import Layout from './components/Layout';
import {
  ErrorPage,
  FrontPage,
  NewcommentsPage,
  NewestPage,
  NewsAskPage,
  NewsfaqPage,
  NewsguidelinesPage,
  NewsItemPage,
  NewsPage,
  NewsShowPage,
  SecurityPage
} from './pages';
import { NewsJobsPage } from './pages/NewsJobs';

interface RouteConfig {
  path: string;
  element: ReactNode;
  children?: RouteConfig[];
}

const routes: RouteConfig[] = [
  {
    path: '/',
    element: <Layout />,
    children: [
      { path: '/', element: <NewsPage /> },
      { path: '/news', element: <NewsPage /> },
      { path: '/newest', element: <NewestPage /> },
      { path: '/front', element: <FrontPage /> },
      { path: '/item', element: <NewsItemPage /> },
      { path: '/newcomments', element: <NewcommentsPage /> },
      { path: '/ask', element: <NewsAskPage /> },
      { path: '/show', element: <NewsShowPage /> },
      { path: '/jobs', element: <NewsJobsPage /> },
      { path: '/submit', element: <NewsPage /> },
      { path: '/error', element: <ErrorPage /> },
      { path: '*', element: <ErrorPage /> }
    ]
  },
  { path: '/newsguidelines', element: <NewsguidelinesPage /> },
  { path: '/newsfaq', element: <NewsfaqPage /> },
  { path: '/security', element: <SecurityPage /> }
];

const App: FC = () => {
  return (
    <Router
      future={{
        v7_startTransition: true,
        v7_relativeSplatPath: true
      }}>
      <Routes>
        {routes.map((route, index) => (
          <Route key={index} path={route.path} element={route.element}>
            {route.children?.map((child, childIndex) => (
              <Route key={childIndex} path={child.path} element={child.element} />
            ))}
          </Route>
        ))}
      </Routes>
    </Router>
  );
};

export default App;
