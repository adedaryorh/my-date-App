import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

class ItemComponent {
  ItemComponent({required this.assetPath, required this.title});
  final String assetPath;
  final String title;
}

final List<ItemComponent> profileItemComponents = [
  ItemComponent(assetPath: AppAssets.profileFriends, title: 'Friends'),
  ItemComponent(assetPath: AppAssets.businessIcon, title: 'Business'),
  ItemComponent(assetPath: AppAssets.contentsIcon, title: 'Contents'),
  ItemComponent(assetPath: AppAssets.celebratedIcon, title: 'Celebrated'),
];
final List<ItemComponent> profileBusinessItemComponents = [
  ItemComponent(assetPath: AppAssets.profileFriends, title: 'Followers'),
  ItemComponent(assetPath: AppAssets.businessIcon, title: 'Business'),
  ItemComponent(assetPath: AppAssets.contentsIcon, title: 'Contents'),
  ItemComponent(assetPath: AppAssets.celebratedIcon, title: 'Celebrated'),
];

class ProfileItem extends StatelessWidget {
  const ProfileItem({
    required this.model,
    super.key,
  });

  final ItemComponent model;

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 68,
      width: double.maxFinite,
      margin: const EdgeInsets.only(left: 20, right: 20, bottom: 10),
      padding: const EdgeInsets.only(left: 20, right: 30),
      decoration: const BoxDecoration(
        color: Color(0xffF4F4F4),
        borderRadius: BorderRadius.all(Radius.circular(10)),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Row(
            children: [
              SvgPicture.asset(
                model.assetPath,
                width: 19,
                height: 19,
              ),
              const Space(30),
              Text(model.title),
            ],
          ),
          const Icon(
            Icons.arrow_forward_ios,
            size: 11,
          ),
        ],
      ),
    );
  }
}
