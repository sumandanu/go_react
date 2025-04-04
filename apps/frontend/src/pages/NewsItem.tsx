import { FC } from 'react';
import { Link, useSearchParams } from 'react-router-dom';

import { fetcher, useApi } from '../hooks/useApi';
import { IComment } from '../interface/IComment';

export const NewsItemPage: FC = () => {
  const [searchParams] = useSearchParams();

  // Get a single parameter
  const id = searchParams.get('id');

  const { data: response, error, isLoading } = useApi<IComment[]>(`/item?id=${id}`, fetcher);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading: {error.message}</div>;

  const news = response?.data || [];
  const submission: any = response?.submission || {};

  return (
    <>
      <ol style={{ display: 'flex', flexDirection: 'column' }}>
        <li style={{ marginBottom: '10px', listStyleType: 'none' }}>
          <div style={{ display: 'flex', flexDirection: 'column', color: 'graytext' }}>
            <div style={{ display: 'inline-flex', gap: 10, alignItems: 'end' }}>
              <Link to={submission.titleUrl} style={{ color: '#000000', textDecoration: 'none' }}>
                <span>{submission.title}</span>
              </Link>
              <Link to={submission.siteUrl} style={{ textDecoration: 'none' }}>
                <span style={{ fontSize: 12, color: '#828282' }}>( {submission.site} )</span>
              </Link>
            </div>
            <div
              style={{
                display: 'inline-flex',
                gap: 4,
                color: 'GrayText',
                fontSize: 12,
                alignItems: 'center'
              }}>
              <span>{submission.score}</span>
              <span>by</span>
              <Link
                to={'/' + submission.byUrl}
                style={{ color: 'graytext', textDecoration: 'none' }}>
                <span>{submission.byUser}</span>
              </Link>
              <Link
                to={'/' + submission.ageUrl}
                style={{ color: 'graytext', textDecoration: 'none' }}>
                <span>{submission.age}</span>
              </Link>
              <span>|</span>
              <Link
                to={'/' + submission.hideUrl}
                style={{ color: 'graytext', textDecoration: 'none' }}>
                <span>hide</span>
              </Link>
              <span>|</span>
              <Link
                to={'/' + submission.pastUrl}
                style={{ color: 'graytext ', textDecoration: 'none' }}>
                <span>past</span>
              </Link>
              <span>|</span>
              <Link
                to={'/' + submission.favoriteUrl}
                style={{ color: 'graytext ', textDecoration: 'none' }}>
                <span>favorite</span>
              </Link>
              <span>|</span>
              <Link
                to={'/' + submission.commentsUrl}
                style={{ color: 'graytext', textDecoration: 'none' }}>
                <span>{submission.comments}</span>
              </Link>
            </div>
            <form
              action="comment"
              method="post"
              style={{ marginBottom: '20px', marginTop: '20px' }}>
              <input type="hidden" name="parent" value="43559605" />
              <input type="hidden" name="goto" value="item?id=43559605" />
              <input type="hidden" name="hmac" value="7e9b276cc28900246e38d2ac0ffc733533f9ae2f" />
              <textarea name="text" rows={8} cols={80} wrap="virtual"></textarea>
              <br />
              <br />
              <input type="submit" value="add comment" />
            </form>
          </div>
        </li>

        {news?.map((list) => (
          <li key={list.id} style={{ marginBottom: '10px', listStyleType: 'none' }}>
            <div style={{ display: 'flex', flexDirection: 'column' }}>
              <div
                style={{
                  display: 'inline-flex',
                  gap: 4,
                  color: 'GrayText',
                  fontSize: 12,
                  alignItems: 'center'
                }}>
                <span>{list.byUser}</span>
                <span>{list.age}</span>
                <span>|</span>
                <Link to={'/' + list.nextUrl} style={{ color: 'graytext', textDecoration: 'none' }}>
                  <span>next</span>
                </Link>
              </div>
              <div
                style={{
                  display: 'inline-flex',
                  gap: 10,
                  alignItems: 'end',
                  fontFamily: 'Verdana, Geneva, sans-serif',
                  fontSize: '9pt'
                }}>
                {list.comment}
              </div>
            </div>
          </li>
        ))}
      </ol>
    </>
  );
};
