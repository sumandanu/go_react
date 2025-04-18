import { FC, useState } from 'react';

import { fetcher, useApi } from '../hooks/useApi';
import { INews } from '../interface/INews';
import ListNews from '../components/ListNews';

export const NewsJobsPage: FC = () => {
  const [pageUrl, setPageUrl] = useState<string | ''>('');

  const { data: response, error, isLoading } = useApi<INews[]>(`/jobs${pageUrl}`, fetcher);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading: {error.message}</div>;

  return (
    <>
      <div style={{ marginLeft: '14px', marginTop: '6px', marginBottom: '14px' }}>
        These are jobs at YC startups. See more at{' '}
        <a href="https://www.ycombinator.com/jobs">
          <u>ycombinator.com/jobs</u>
        </a>
        .
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
