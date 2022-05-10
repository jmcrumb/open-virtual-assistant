export class Review {
    id: string;
    source_link: string;
    account: string;
    plugin: string;
    rating: number;
    content: string;

    constructor(json: { [key: string]: any }) {
        this.id = json.id;
        this.source_link = json.source_link;
        this.account = json.account;
        this.plugin = json.plugin;
        this.rating = json.rating;
        this.content = json.content;
    }

    public static average(reviews: Review[]) {
        let val = 0.0;

        reviews.forEach((review) => {
            val += review.rating;
        });

        return val / reviews.length;
    }
}

export class Report {
    id: string;
    user: string;
    plugin: string;
    content: string;
    is_resolved: string;

    constructor(json: { [key: string]: any }) {
        this.id = json.id;
        this.user = json.user;
        this.plugin = json.plugin;
        this.content = json.content;
        this.is_resolved = json.isResolved;
    }
}

export class Plugin {
    id: string;
    name: string;
    publisher: string;
    sourceLink: string;
    about: string;
    download_count: number;
    published_on: Date;
    reviews: Review[];

    constructor(json: { [key: string]: any }) {
        this.id = json.id;
        this.name = json.name;
        this.publisher = json.publisher;
        this.sourceLink = json.sourceLink;
        this.about = json.about;
        this.download_count = json.download_count;
        this.published_on = json.published_on;
        this.reviews = [];
    }
}