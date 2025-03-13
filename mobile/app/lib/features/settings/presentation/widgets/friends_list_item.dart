import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class FriendsListItem extends StatefulWidget {
  const FriendsListItem({
    super.key,
  });

  @override
  State<FriendsListItem> createState() => _FriendsListItemState();
}

class _FriendsListItemState extends State<FriendsListItem> {
  bool select = false;
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 10),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Row(
            children: [
              Container(
                height: 48,
                width: 48,
                decoration: const ShapeDecoration(
                  shape: CircleBorder(),
                  color: Color(0xffC2C9D6),
                ),
              ),
              const SizedBox(
                width: 10,
              ),
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    'tinatopaz',
                    style: context.textTheme.bodyMedium
                        ?.copyWith(fontWeight: FontWeight.w700),
                  ),
                  const SizedBox(
                    height: 5,
                  ),
                  Text(
                    'Tina Topaz',
                    style: context.textTheme.bodyLarge
                        ?.copyWith(fontWeight: FontWeight.w500),
                  ),
                ],
              ),
            ],
          ),
          Transform.scale(
            scale: 0.7,
            child: Checkbox(
              materialTapTargetSize: MaterialTapTargetSize.shrinkWrap,
              value: select,
              onChanged: (val) {
                if (val == null) return;
                setState(() {
                  select = val;
                });
              },
            ),
          ),
        ],
      ),
    );
  }
}
