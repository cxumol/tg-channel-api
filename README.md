Telegram Channel API
========================

The JSON web API for public Telegram channels. 

Try it now:

- Check what Durov is up to: https://t-me.vercel.app/durov
- Check the second-latest Durov's post: https://t-me.vercel.app/durov/2

Usage
-----

You can use the Web API to get the latest messages, or any one of recent 20 messages, from any public Telegram channel.

The official demo is at `t-me.vercel.app`, and this API is available as `GET /:channel_username/:message_id?`

parameters:
- `channel_username` (string) : the substring of the channel link after "t.me", e.g. `durov` is the `channel_username` of `t.me/durov` 
- `message_id` (integer, within [1, 20] for now) : the ordinal number of the message, e.g. `1` for the latest message, `2` for the second-latest message, etc.

Technical details
-----------------

It's running on [Vercel](https://vercel.com/) and uses its [Go runtime](https://vercel.com/docs/runtimes#official-runtimes/go) to serve the API.

To deploy your own instance, simplely for this repository and connect your copy to [Vercel](https://vercel.com/). 


TO DO
-----

- [ ] Add support for optional fields
  + photos `tgme_widget_message_photo_wrap`
  + sender's name `tgme_widget_message_from_author` 
  + etc.

Contributing
------------

Contributions are welcome!

Try these ideas:
- make progress to the to-do list above
- deploy to a new platform
- add a new feature
- fix a bug
- improve the performance
