import { FC, useState } from 'react';

import { fetcher, useApi } from '../hooks/useApi';
import { IComment } from '../interface/IComment';
import ListNews from '../components/ListNews';

export const NewcommentsPage: FC = () => {
  const [pageUrl, setPageUrl] = useState<string | ''>('');

  const {
    data: response,
    error,
    isLoading
  } = useApi<IComment[]>(`/newcomments${pageUrl}`, fetcher);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading: {error.message}</div>;

  return (
    <ListNews
      type="comment"
      data={response?.data || []}
      showMore
      onMoreClick={() =>
        setPageUrl(response?.nextLink ? `?${response?.nextLink?.split('?')[1]}` : '')
      }
    />
  );
};
