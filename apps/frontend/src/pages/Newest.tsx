import { FC, useState } from 'react';

import { fetcher, useApi } from '../hooks/useApi';
import { INews } from '../interface/INews';
import ListNews from '../components/ListNews';

export const NewestPage: FC = () => {
  const [pageUrl, setPageUrl] = useState<string | ''>('');

  const { data: response, error, isLoading } = useApi<INews[]>(`/newest${pageUrl}`, fetcher);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading: {error.message}</div>;

  return (
    <ListNews
      data={response?.data || []}
      showMore
      onMoreClick={() =>
        setPageUrl(response?.nextLink ? `?${response?.nextLink?.split('?')[1]}` : '')
      }
    />
  );
};
