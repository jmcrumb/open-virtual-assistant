
import { Plugin, Review, useQueryPluginByID } from "../api/pluginStoreAPI";
import * as React from "react";
import { useQueryClient } from "react-query";
import Container from "@mui/material/Container";
import { Box, Button, Card, CardActions, CardContent, Chip, Grid, Modal, Paper, Rating, Skeleton, Stack, Typography } from "@mui/material";
import ReactTimeAgo from 'react-time-ago';
import CodeIcon from '@mui/icons-material/Code';
import UpdateIcon from '@mui/icons-material/Update';
import EditIcon from '@mui/icons-material/Edit';
import { height } from "@mui/material/node_modules/@mui/system";


export function PluginViewPublic(props) {
  const styledElevation: number = 2;

  // Modal controls
  const [open, setOpen] = React.useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);

  // const queryClient = useQueryClient();
  // const { status, data, error, isFetching } = useQueryPluginByID(
  //   props.id
  // );

  // const plugin: Plugin = data;
  // let reviews: Review[] = plugin.getReviews();
  // let rating: number = Review.average(reviews);

  const plugin: Plugin = new Plugin({
    "id": "3f094753-6d45-4897-a749-c51378ddbe13",
    "name": "test plugin",
    "publisher": "29b744e6-0a2f-48a9-aeb0-1a162f546763",
    "source_link": "http://127.0.0.1:8000",
    "about": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
    "download_count": 10000,
    "published_on": "2022-04-29T00:00:00Z"
  });
  let reviews: Review[] = [
    {
      "id": "d0c96e62-2191-44d7-9203-458293b43071",
      "source_review": "",
      "account": "29b744e6-0a2f-48a9-aeb0-1a162f546763",
      "plugin": "3f094753-6d45-4897-a749-c51378ddbe13",
      "rating": 4.5,
      "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
    },
    {
      "id": "1bc50705-0c1e-4150-a01b-5c60d82cc8a7",
      "source_review": "",
      "account": "29b744e6-0a2f-48a9-aeb0-1a162f546763",
      "plugin": "3f094753-6d45-4897-a749-c51378ddbe13",
      "rating": 2,
      "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."
    }
  ];
  let status = "success";

  let rating: number = Review.average(reviews);

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
                fontSize: '2rem'
              }}>
                Install
              </Button>
            </Grid>
            <Grid item>
              <Stack direction="row" spacing={1} id="plugin-quick-info">
                <PluginRating value={rating} />
                <Chip
                  variant="outlined"
                  label={formatDownloadCount(plugin.download_count) + " users"}
                />
                <Chip
                  variant="outlined"
                  label={<ReactTimeAgo date={new Date(plugin.published_on)} locale="en-US" />}
                  icon={<UpdateIcon />}
                />
                <div>
                  <Chip
                    variant="outlined"
                    label="Source"
                    onClick={handleOpen}
                    icon={<CodeIcon />}
                    color="primary"
                  />
                  <Modal
                    open={open}
                    onClose={handleClose}
                    aria-labelledby="modal-modal-title"
                    aria-describedby="modal-modal-description"
                  >
                    <Box sx={{
                      position: 'absolute' as 'absolute',
                      top: '50%',
                      left: '50%',
                      transform: 'translate(-50%, -50%)',
                      width: 400,
                      bgcolor: 'background.paper',
                      border: '2px solid #000',
                      boxShadow: 24,
                      p: 4,
                    }}>
                      <Typography id="modal-modal-title" variant="h6" component="h2">
                        This link will take you to the outside source below.  Click continue to proceed
                      </Typography>
                      <Typography id="modal-modal-description" sx={{ mt: 2 }}>
                        {plugin.sourceLink}
                      </Typography>
                    </Box>
                  </Modal>
                </div>


              </Stack>
            </Grid>
          </Grid>
          <Grid
            container
            spacing={0}
            direction="column"
          >
            <Grid item sx={{ paddingBottom: '2rem' }}>
              <Typography variant="h3" sx={{ alignItems: 'left' }}>Description</Typography>
              <Paper elevation={styledElevation} className="padded-paper">
                <p>{plugin.about}</p>
              </Paper>
            </Grid>
            <Grid item>
              <Box
                sx={{
                  height: 'auto'
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
                    fontSize: '1.5rem'
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
          >
            <Grid item>
              <Paper elevation={styledElevation} className="padded-paper">
                <ul>
                  {reviews.map((review) => <ReviewCard data={review} />)}
                </ul>
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
    <li key={review.id}>
      <Card>
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
    </li>
  );
}

function PluginRating(props) {
  return (
    <Rating name="half-rating-read" defaultValue={props.value} precision={0.5} readOnly />
  )
}
