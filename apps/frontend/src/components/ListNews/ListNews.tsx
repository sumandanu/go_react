import { Link } from 'react-router-dom';

import { INews } from '../../interface/INews';
import styles from './Styles.module.css';
import Item from './Item';
import ItemComment from './ItemComment';
import { IComment } from '../../interface/IComment';

interface Props {
  type?: 'comment' | 'news' | undefined;
  data: INews[] | IComment[];
  nextLink?: string;
}

const ListNews: React.FC<Props> = (props) => {
  return (
    <>
      <ol>
        {props.data.map((list) =>
          props?.type === 'comment' ? (
            <ItemComment {...(list as IComment)} />
          ) : (
            <Item {...(list as INews)} />
          )
        )}
      </ol>
      <Link to={`/${props?.nextLink}`} className={styles.more}>
        More
      </Link>
    </>
  );
};

export default ListNews;
