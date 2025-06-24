import 'package:flutter/material.dart';

class ProfileCard extends StatelessWidget {
  final String name;
  final String email;
  final int age;
  final String? avatarUrl;

  const ProfileCard({
    Key? key,
    required this.name,
    required this.email,
    required this.age,
    this.avatarUrl,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // TODO: Implement profile card UI
    return Card(  // this widget serves as a container for related content and actions on the same topic
      elevation: 20,
      child: Padding(
        padding: EdgeInsets.all(50),
        child: Column(
          crossAxisAlignment:
            CrossAxisAlignment.center,
          children: [
            CircleAvatar(
              radius: 25,
              child: (avatarUrl == null || avatarUrl!.isEmpty) ? Text(name[0]) : null,
            ),
            SizedBox(height: 20),
            Text(
              name,
              style: TextStyle(fontSize: 20, fontWeight: FontWeight.w900), 
            ),
            SizedBox(height: 10),
            Text(
              email,
              style: TextStyle(fontSize: 15, fontWeight: FontWeight.w900),
            ),
            SizedBox(height: 10),
            Text(
              'Age: $age',
              style: TextStyle(fontSize: 15, fontWeight: FontWeight.w900),
            ),
            
          ],  
        )
      )
    );
  }
}
