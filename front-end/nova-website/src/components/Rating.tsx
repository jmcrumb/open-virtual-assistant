import * as React from 'react';

function Rating(props) {
	let {value} = props

	return (
		<div className="Rating">
			<div className="c-rating c-rating--small" data-rating-value={value}>
				<button>1</button>
				<button>2</button>
				<button>3</button>
				<button>4</button>
				<button>5</button>
			</div>
		</div>
	);
}

export default Rating;