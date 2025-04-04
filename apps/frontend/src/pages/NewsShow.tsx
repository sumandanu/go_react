import { FC } from 'react';

import { fetcher, useApi } from '../hooks/useApi';
import { INews } from '../interface/INews';
import ListNews from '../components/ListNews';

export const NewsShowPage: FC = () => {
  const { data: response, error, isLoading } = useApi<INews[]>(`/show`, fetcher);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading: {error.message}</div>;

  return (
    <>
      <div style={{ marginLeft: '36px', marginTop: '6px', marginBottom: '12px' }}>
        Please read the Show HN{' '}
        <a href="showhn.html">
          <u>rules</u>
        </a>{' '}
        and{' '}
        <a href="item?id=22336638">
          <u>tips</u>
        </a>{' '}
        before posting. You can browse the newest Show HNs{' '}
        <a href="shownew">
          <u>here</u>
        </a>
        .
      </div>
      <ListNews data={response?.data || []} nextLink={response?.nextLink} />
    </>
  );
};
