import { INews } from '../../interface/INews';
import styles from './Styles.module.css';
import Item from './Item';
import ItemComment from './ItemComment';
import { IComment } from '../../interface/IComment';

interface Props {
  type?: 'comment' | 'news' | undefined;
  data: INews[] | IComment[];
  showMore?: boolean;
  onMoreClick?: () => void;
}

const ListNews: React.FC<Props> = (props) => {
  return (
    <div style={{ display: 'flex', flexDirection: 'column' }}>
      <ol>
        {props.data.map((list) =>
          props?.type === 'comment' ? (
            <ItemComment {...(list as IComment)} />
          ) : (
            <Item {...(list as INews)} />
          )
        )}
      </ol>
      <button
        onClick={props?.onMoreClick}
        className={styles.more}
        style={{ display: !props?.showMore ? 'none' : 'block' }}>
        More
      </button>
    </div>
  );
};

export default ListNews;
