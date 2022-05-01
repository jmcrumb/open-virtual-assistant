import * as React from 'react';
import Rating from './Rating';

function PluginPreview(props) {
	let {thumbnail, name, author, rating, description} = props

	return (
		<div className="PluginPreview">
			<img src={thumbnail} alt={name + " plugin thumbnail"} className="thumbnail" />
			<div className="info">
				<span className="name">{name}</span>
				<span className="author">{author}</span>
				<Rating value={rating} />
			</div>
			<div className="desc">
				<p>{description}</p>
			</div>
		</div>
	);
}

export default PluginPreview;