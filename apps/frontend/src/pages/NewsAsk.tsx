import { FC, useState } from 'react';

import { fetcher, useApi } from '../hooks/useApi';
import ListNews from '../components/ListNews';
import { INews } from '../interface/INews';

export const NewsAskPage: FC = () => {
  const [pageUrl, setPageUrl] = useState<string | ''>('');

  const { data: response, error, isLoading } = useApi<INews[]>(`/ask${pageUrl}`, fetcher);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading: {error.message}</div>;

  return (
    <ListNews
      data={response?.data || []}
      onMoreClick={() =>
        setPageUrl(response?.nextLink ? `?${response?.nextLink?.split('?')[1]}` : '')
      }
    />
  );
};
