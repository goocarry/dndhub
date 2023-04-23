# dndhub
Back-end part of D&amp;DHub mobile app.

# Usage
All the make commands should be run from root directory.


# Deploy stack
1. Remove previously deployed stack with
   ```
   docker stack rm dndhub
   ```
2. Update /swarm/swarm.yml file with new from repository
3. Deploy stack with
   ```
   docker stack deploy -c swarm.yml dndhub
   ```

### !!! Warning 
Build docker images is needed to be done from same architecture as production mashine!


