import { FC, useState } from 'react';

import { fetcher, useApi } from '../hooks/useApi';
import { INews } from '../interface/INews';
import ListNews from '../components/ListNews';

export const FrontPage: FC = () => {
  const [pageUrl, setPageUrl] = useState<string | ''>('');

  const { data: response, error, isLoading } = useApi<INews[]>(`/front${pageUrl}`, fetcher);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading users: {error.message}</div>;

  return (
    <>
      <div style={{ marginLeft: '36px', marginTop: '6px', marginBottom: '14px' }}>
        Stories from April 2, 2025 (UTC)
        <div style={{ marginTop: '9px' }}>
          Go back a{' '}
          <span className="hnmore">
            <a href="front?day=2025-04-01">day</a>
          </span>
          ,{' '}
          <span className="hnmore">
            <a href="front?day=2025-03-02">month</a>
          </span>
          , or{' '}
          <span className="hnmore">
            <a href="front?day=2024-04-02">year</a>
          </span>
          . Go forward a{' '}
          <span className="hnmore">
            <a href="front?day=2025-04-03">day</a>
          </span>
          .
        </div>
      </div>
      <ListNews
        data={response?.data || []}
        showMore
        onMoreClick={() =>
          setPageUrl(response?.nextLink ? `?${response?.nextLink?.split('?')[1]}` : '')
        }
      />
    </>
  );
};
