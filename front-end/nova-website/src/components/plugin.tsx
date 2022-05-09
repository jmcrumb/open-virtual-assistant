
import { Plugin, Review } from "../api/pluginStoreAPI";
import * as React from "react";
import Container from "@mui/material/Container";
import { Box, Button, Card, CardActions, CardContent, Chip, Grid, Link, Modal, Paper, Rating, Skeleton, Stack, Typography } from "@mui/material";
import ReactTimeAgo from 'react-time-ago';
import CodeIcon from '@mui/icons-material/Code';
import UpdateIcon from '@mui/icons-material/Update';
import EditIcon from '@mui/icons-material/Edit';
import axios from "axios";
import { BACKEND_SRC } from "../api/helper";
import { useParams } from "react-router-dom";


export function PluginViewPublic() {
  const styledElevation: number = 2;
  const pluginId = useParams();
  // const pluginId= "3f094753-6d45-4897-a749-c51378ddbe13";

  // Plugin controls
  const [plugin, setPlugin] = React.useState(null);
  const [reviews, setReviews] = React.useState([]);

  let rating = null;

  React.useEffect(() => {
    axios.get(`${BACKEND_SRC}plugin/${pluginId}`).then((response) => {
      setPlugin(new Plugin(response.data));
    });

    axios.get(`${BACKEND_SRC}review/${pluginId}`).then((response) => {
      let temp = [];
      response.data.forEach((r: { [key: string]: any; }) => {
        temp.push(new Review(r));
      });
      setReviews(temp);
    });
  }, []);

  if (!plugin) return null;

  const formatDownloadCount = ((count) => {
    if (count >= 1000 && count < 1000000) {
      return `${count / 1000}k`
    } else if (count >= 1000000) {
      return `${count / 1000000}m`
    } else {
      return `${count}`
    }
  })

  return (
    <div>
        <Container>
          <Grid container spacing={2} sx={{ paddingBottom: '2rem' }}>
            <Grid item xs={8}>
              <Typography variant="h1" sx={{ display: 'inline' }}>{plugin.name}</Typography>
              <Button variant="text">
                <Typography variant="h6">Published by John Doe</Typography>
              </Button>
            </Grid>
            <Grid item xs={4}>
              <Button variant="contained" sx={{
                position: 'absolute',
                top: '10%',
                right: '3rem',
                fontSize: '2rem'
              }}>
                Install
              </Button>
            </Grid>
            <Grid item>
              <Stack direction="row" spacing={1} id="plugin-quick-info">
                {reviews == [] ? <Skeleton/> : <PluginRating value={Review.average(reviews)} />}
                
                <Chip
                  variant="outlined"
                  label={formatDownloadCount(plugin.download_count) + " users"}
                />
                <Chip
                  variant="outlined"
                  label={<ReactTimeAgo date={new Date(plugin.published_on)} locale="en-US" />}
                  icon={<UpdateIcon />}
                />
                <Link href={plugin.source_link} target="_blank" rel="noopener">
                {<Chip
                    variant="outlined"
                    label="Source"
                    icon={<CodeIcon />}
                    color="primary"
                  />}
                </Link>
              </Stack>
            </Grid>
          </Grid>
          <Grid
            container
            spacing={0}
            direction="column"
          >
            <Grid item>
              <Typography variant="h3" sx={{ 
                alignItems: 'left', 
                marginBottom: '1.5rem' 
              }}>
                  Description
              </Typography>
              <Paper elevation={styledElevation} className="padded-paper">
                <p>{plugin.about}</p>
              </Paper>
            </Grid>
            <Grid item>
              <Box
                sx={{
                  height: 'auto',
                  m: '1.5rem'
                }}
              >
                <Typography
                  variant="h3"
                  sx={{
                    display: 'inline',
                    marginTop: '0.25rem',
                    marginBottom: '0.25rem'
                  }}
                >
                  Reviews</Typography>
                <Button
                  variant="contained"
                  endIcon={<EditIcon />}
                  sx={{
                    position: 'absolute',
                    right: 0,
                    fontSize: '1.5rem',
                    marginRight: '3rem'
                  }}
                >
                  Write Review
                </Button>
              </Box>
            </Grid>
          </Grid>
          <Grid
            container
            spacing={0}
            direction="column"
            alignItems="center"
            justifyContent="center"
            sx={{
              width: 'inherit'
            }}
          >
            <Grid item>
              <Paper elevation={styledElevation} className="padded-paper">
                <Stack direction="column" spacing={2}>
                  {reviews.map((review) => <ReviewCard data={review} />)}
                </Stack>
              </Paper>
            </Grid>
          </Grid>
        </Container>

    </div>
  );
}

function ReviewCard(props) {
  const review: Review = props.data;
  return (
      <Card sx={{width: `${screen.availWidth}px`}}>
      <CardContent>
        {/* TODO: API call for account name */}
        <Typography variant="h5">John Doe</Typography>
        <PluginRating value={review.rating} />
        <p>{review.content}</p>
      </CardContent>
      <CardActions>
        <Button sx={{ fontSize: '22px' }}>Comment</Button>
      </CardActions>
    </Card>
  );
}

function PluginRating(props) {
  return (
    <Rating name="half-rating-read" defaultValue={props.value} precision={0.5} readOnly />
  )
}
