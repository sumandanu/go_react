import { Link } from 'react-router-dom';
import styles from './Styles.module.css';
import { IComment } from '../../interface/IComment';

const ItemComment: React.FC<IComment> = (props) => {
  return (
    <li key={props.id} style={{ listStyleType: 'none' }}>
      <div className={styles.card}>
        <Link to={`/${props.voteUrl}`}>
          <div className={styles.votearrow} title="upvote"></div>
        </Link>
        <div className={styles.content}>
          <div className={styles.subTitle}>
            <span className={styles.siteText}>by</span>
            <Link to={`/${props.byUrl}`} className={!props.byUrl ? styles.hidden : ''}>
              <span className={styles.siteText}>{props.byUser}</span>
            </Link>
            <Link to={`/${props.ageUrl}`} className={!props.ageUrl ? styles.hidden : ''}>
              <span className={styles.siteText}>{props.age}</span>
            </Link>
            <span className={`${styles.siteText} ${!props.parentUrl ? styles.hidden : ''}`}>|</span>
            <Link to={`/${props.parentUrl}`} className={!props.parentUrl ? styles.hidden : ''}>
              <span className={styles.siteText}>parent</span>
            </Link>
            <span className={`${styles.siteText} ${!props.contextUrl ? styles.hidden : ''}`}>
              |
            </span>
            <Link to={`/${props.contextUrl}`} className={!props.contextUrl ? styles.hidden : ''}>
              <span className={styles.siteText}>context</span>
            </Link>
            <span className={`${styles.siteText} ${!props.onStoryUrl ? styles.hidden : ''}`}>
              | on:{' '}
            </span>
            <Link to={`/${props.onStoryUrl}`} className={!props.onStoryUrl ? styles.hidden : ''}>
              <span className={styles.siteText}>{props.onStory}</span>
            </Link>
          </div>
          <div className={styles.title} style={{ marginRight: '60px' }}>
            <span className={styles.titleText}>{props.comment}</span>
          </div>
        </div>
      </div>
    </li>
  );
};

export default ItemComment;
